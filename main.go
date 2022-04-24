package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/sokdak/miner-exporter/pkg/common"
	"github.com/sokdak/miner-exporter/pkg/metric"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	var port int
	var minerType string
	var host string
	var protocol string
	var listenPort int
	flag.StringVar(&minerType, "miner-type", "gminer", "Type of miner")
	flag.StringVar(&host, "host", "localhost", "host for export miner metric")
	flag.StringVar(&protocol, "protocol", "http", "protocol for export miner metric")
	flag.IntVar(&port, "port", 8080, "Port that retrieve status")
	flag.IntVar(&listenPort, "listen-port", 12000, "port for listening")
	flag.Parse()

	glog := metric.GetLoggerOrDie()
	log := glog.WithName("main")
	log.Info("starting miner-exporter")

	var err error
	for retry := 0; retry < common.InitRetryThreshold; retry++ {
		err = metric.SetMinerInstanceOrDie(minerType, protocol, host, port, glog)
		if err != nil {
			glog.Error(err, "SetMinerInstance has been failed, retrying..", "retry", retry, "retry-threshold", common.InitRetryThreshold)
			time.Sleep(common.InitRetryBackoff)
		} else {
			break
		}
	}

	if err != nil {
		glog.Error(err, "failed get miner instance, exit miner.")
		os.Exit(1)
	}

	metric.SetLoggerForMetricHandler(glog)

	// create a http server
	server := &http.Server{}

	// Create a context that is cancelled on SIGKILL or SIGINT.
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	mux := http.DefaultServeMux
	mux.Handle("/metrics", new(metric.JsonHandler))
	server.Handler = mux
	server.Addr = fmt.Sprintf(":%d", listenPort)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Error(err, "failed to listen port, metric endpoint won't be exposed.")
			os.Exit(1)
		}
	}()
	<-ctx.Done()
	log.Info("shutting down")

	// create a context for graceful http server shutdown
	srvCtx, srvCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer srvCancel()

	_ = server.Shutdown(srvCtx)
}
