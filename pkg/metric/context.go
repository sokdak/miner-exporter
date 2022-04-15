package metric

import (
	"github.com/gin-gonic/gin"
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"github.com/sokdak/miner-exporter/pkg/common"
	"go.uber.org/zap"
	"net/http"
)

func GetMiner(c *gin.Context) common.MinerInterface {
	if _miner != nil {
		return _miner
	}
	c.String(http.StatusInternalServerError, "failed to get miner instance")
	c.Abort()
	return nil
}

func GetLogger(c *gin.Context) logr.Logger {
	if _logger.Enabled() {
		return _logger
	}
	c.String(http.StatusInternalServerError, "failed to get logger")
	c.Abort()
	return zapr.NewLogger(zap.NewNop())
}
