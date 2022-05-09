package metric

import (
	"fmt"
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	cgminer "github.com/sokdak/go-teamredminer-api"
	"github.com/sokdak/miner-exporter/pkg/common"
	"github.com/sokdak/miner-exporter/pkg/gminer"
	"github.com/sokdak/miner-exporter/pkg/nbminer"
	"github.com/sokdak/miner-exporter/pkg/teamredminer"
	"github.com/sokdak/miner-exporter/pkg/trex"
	"go.uber.org/zap"
	"net/http"
	"os"
)

var (
	_miner  common.MinerInterface
	_logger logr.Logger
)

func GetLoggerOrDie() logr.Logger {
	z, err := zap.NewDevelopment()
	if err != nil {
		fmt.Println("failed to init logger")
		os.Exit(1)
	}
	return zapr.NewLogger(z)
}

func SetLoggerForMetricHandler(log logr.Logger) {
	_logger = log.WithName("exporter")
}

func SetMinerInstanceOrDie(minerType string, protocol string, host string, port int, log logr.Logger) error {
	logInit := log.WithName("init")
	connInfo := common.ConnectionInfo{
		Protocol: protocol,
		Host:     host,
		Port:     port,
	}

	var miner common.MinerInterface
	switch minerType {
	case common.MinerTypeGMiner:
		miner = gminer.Client{
			Log:            log.WithName("gminer"),
			ConnectionInfo: connInfo,
			HttpClient: http.Client{
				Timeout: ScrapeTimeout,
			},
		}
	case common.MinerTypeTrexMiner:
		miner = trex.Client{
			Log:            log.WithName("trex"),
			ConnectionInfo: connInfo,
			HttpClient: http.Client{
				Timeout: ScrapeTimeout,
			},
		}
	case common.MinerTypeTeamRedMiner:
		miner = teamredminer.Client{
			Log:            log.WithName("teamredminer"),
			ConnectionInfo: connInfo,
			Client:         cgminer.NewCGMiner(connInfo.Host, connInfo.Port, ScrapeTimeout),
		}
	case common.MinerTypeNbMiner:
		miner = nbminer.Client{
			Log:            log.WithName("nbminer"),
			ConnectionInfo: connInfo,
			HttpClient: http.Client{
				Timeout: ScrapeTimeout,
			},
		}
	default:
		logInit.WithValues("miner-type", minerType,
			"connection-url", common.GetConnectionString(connInfo)).Error(nil, "not supported miner type")
		os.Exit(1)
	}

	if err := miner.Init(); err != nil {
		logInit.WithValues("miner-type", minerType,
			"connection-url", common.GetConnectionString(connInfo)).Error(err, "failed to init miner")
		return err
	}
	if err := miner.Ping(); err != nil {
		logInit.WithValues("miner-type", minerType,
			"connection-url", common.GetConnectionString(connInfo)).Error(err, "failed to ping miner")
		return err
	}

	logInit.Info("successfully initialized", "connection-url", common.GetConnectionString(connInfo))
	_miner = miner
	return nil
}

func GetAndSetupHandler(format string, log logr.Logger) (http.Handler, error) {
	var h http.Handler
	switch format {
	case "json":
		h = new(JsonHandler)
	case "prometheus":
		// register metrics to prometheus client
		prometheus.MustRegister(NewMinerMetric(log))
		h = promhttp.Handler()
	default:
		return nil, fmt.Errorf("failed to get handler, handler not supported: %s", format)
	}
	return h, nil
}
