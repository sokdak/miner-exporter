package teamredminer

import (
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"github.com/sokdak/go-teamredminer-api"
	"github.com/sokdak/miner-exporter/pkg/common"
)

type Client struct {
	Log            logr.Logger
	ConnectionInfo common.ConnectionInfo
	Client         *cgminer.CGMiner
}

func (c Client) Ping() error {
	_, err := c.Client.Summary()
	if err != nil {
		return errors.Wrapf(err, "failed to ping teamredminer")
	}
	return nil
}

func (c Client) Init() error {
	return nil
}
