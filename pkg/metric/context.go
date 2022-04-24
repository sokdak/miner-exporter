package metric

import (
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"github.com/sokdak/miner-exporter/pkg/common"
	"go.uber.org/zap"
)

func GetMiner() common.MinerInterface {
	if _miner != nil {
		return _miner
	}
	return nil
}

func GetLogger() logr.Logger {
	if _logger.Enabled() {
		return _logger
	}
	return zapr.NewLogger(zap.NewNop())
}
