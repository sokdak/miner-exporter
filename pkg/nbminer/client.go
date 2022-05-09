package nbminer

import (
	"fmt"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"github.com/sokdak/miner-exporter/pkg/common"
	"net/http"
)

const (
	MinerHttpStatusUrl = "/api/v1/status"
)

type Client struct {
	Log            logr.Logger
	ConnectionInfo common.ConnectionInfo
	HttpClient     http.Client
}

func (c Client) Ping() error {
	return nil
}

func (c Client) Init() error {
	// verify connection info
	_, err := c.HttpClient.Head(fmt.Sprintf("%s%s", common.GetConnectionString(c.ConnectionInfo), MinerHttpStatusUrl))
	if err != nil {
		return errors.Wrapf(err, "failed to get response from endpoint")
	}
	return nil
}
