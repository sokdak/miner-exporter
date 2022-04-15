package trex

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sokdak/miner-exporter/pkg/common"
	"github.com/sokdak/miner-exporter/pkg/dto"
	"io/ioutil"
	"net/http"
)

func (c Client) Fetch() (interface{}, error) {
	resp, err := c.HttpClient.Get(fmt.Sprintf("%s%s", common.GetConnectionString(c.ConnectionInfo), MinerHttpSummaryUrl))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to fetch response from endpoint, is host alive?")
	} else if resp.StatusCode != http.StatusOK {
		c.Log.Info("host returned non-200 status", "url", resp.Request.URL, "status", resp.Status)
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read body")
	}

	summary := &Summary{}
	if err := json.Unmarshal(data, &summary); err != nil {
		return nil, err
	}
	return *summary, nil
}

func (c Client) Parse(value interface{}) (*dto.Status, error) {
	summary, ok := value.(Summary)
	if !ok {
		return nil, fmt.Errorf("failed to parse trex response dto")
	}

	commonStatus := dto.Status{
		Miner: dto.Miner{
			Name:      summary.Name,
			Version:   summary.Version,
			Algorithm: summary.Algorithm,
			Address:   summary.ActivePool.User,
			Pool:      summary.ActivePool.Url,
			Uptime:    summary.Uptime,
		},
		Devices: func() []dto.Device {
			devs := make([]dto.Device, 0)
			for _, dev := range summary.Gpus {
				devs = append(devs, dto.Device{
					GpuId:            dev.DeviceId,
					Name:             common.GeneralizeGpuName(dev.Name),
					Hashrate:         dev.Hashrate,
					FanSpeed:         dev.FanSpeed,
					CoreTemp:         dev.Temperature,
					MemoryTemp:       dev.MemoryTemperature,
					PowerConsumption: dev.Power,
					ShareAccepted:    dev.Shares.AcceptedCount,
					ShareRejected:    dev.Shares.RejectedCount,
					ShareStale:       dev.Shares.InvalidCount,
				})
			}
			return devs
		}(),
	}

	return &commonStatus, nil
}
