package gminer

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sokdak/miner-exporter/pkg/common"
	"github.com/sokdak/miner-exporter/pkg/dto"
	"io/ioutil"
	"net/http"
	"strings"
)

func (c Client) Fetch() (interface{}, error) {
	resp, err := c.HttpClient.Get(fmt.Sprintf("%s%s", common.GetConnectionString(c.ConnectionInfo), MinerHttpStatUrl))
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

	status := &Stat{}
	if err := json.Unmarshal(data, &status); err != nil {
		return nil, err
	}
	return *status, nil
}

func (c Client) Parse(value interface{}) (*dto.Status, error) {
	status, ok := value.(Stat)
	if !ok {
		return nil, fmt.Errorf("failed to parse gminer response dto")
	}

	commonStatus := dto.Status{
		Miner: dto.Miner{
			Name:      getName(status.Miner),
			Version:   getVersion(status.Miner),
			Algorithm: common.GeneralizeAlgorithm(status.Algorithm),
			Address:   common.ExtractAddress(status.User),
			Pool:      common.GeneralizePoolAddress(status.Server),
			Uptime:    status.Uptime,
			Worker:    common.ExtractWorkerNameFromAddress(status.User),
		},
		Devices: func() []dto.Device {
			devices := make([]dto.Device, 0)
			for _, dev := range status.Devices {
				devices = append(devices, dto.Device{
					GpuId:            dev.GpuId,
					Name:             dev.Name,
					Hashrate:         dev.Speed,
					FanSpeed:         dev.Fan,
					CoreTemp:         dev.Temperature,
					MemoryTemp:       dev.MemoryTemperature,
					PowerConsumption: dev.PowerUsage,
					ShareAccepted:    dev.AcceptedShares,
					ShareRejected:    dev.RejectedShares,
					ShareStale:       dev.StaleShares,
				})
			}
			return devices
		}(),
	}
	return &commonStatus, nil
}

func getVersion(a string) string {
	return strings.Split(a, " ")[1]
}

func getName(a string) string {
	return strings.Split(a, " ")[0]
}
