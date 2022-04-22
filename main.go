package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/sokdak/miner-exporter/pkg/common"
	"github.com/sokdak/miner-exporter/pkg/metric"
	"os"
	"time"
)

func main() {
	var port int
	var minerType string
	var host string
	var protocol string
	flag.StringVar(&minerType, "miner-type", "gminer", "Type of miner")
	flag.StringVar(&host, "host", "localhost", "host for export miner metric")
	flag.StringVar(&protocol, "protocol", "http", "protocol for export miner metric")
	flag.IntVar(&port, "port", 8080, "Port that retrieve status")
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
		}
	}

	if err != nil {
		glog.Error(err, "failed get miner instance, exit miner.")
		os.Exit(1)
	}

	metric.SetLoggerForMetricHandler(glog)
	gin.SetMode(gin.ReleaseMode)
	g := gin.Default()
	g.GET("/metrics", metric.HandleExportMetric)
	if err := g.Run(":12000"); err != nil {
		log.Error(err, "failed to run gin")
		os.Exit(1)
	}
}
