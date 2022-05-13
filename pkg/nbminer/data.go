package nbminer

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sokdak/miner-exporter/pkg/common"
	"github.com/sokdak/miner-exporter/pkg/dto"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func (c Client) Fetch() (interface{}, error) {
	resp, err := c.HttpClient.Get(fmt.Sprintf("%s%s", common.GetConnectionString(c.ConnectionInfo), MinerHttpStatusUrl))
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

	status := &Status{}
	if err := json.Unmarshal(data, &status); err != nil {
		return nil, err
	}
	return *status, nil
}

func (c Client) Parse(value interface{}) (*dto.Status, error) {
	status, ok := value.(Status)
	if !ok {
		return nil, fmt.Errorf("failed to parse nbminer response dto")
	}

	commonStatus := dto.Status{
		Miner: dto.Miner{
			Name:      "nbminer",
			Version:   status.Version,
			Algorithm: common.GeneralizeAlgorithm(status.Stratum.Algorithm),
			Address:   common.ExtractAddress(status.Stratum.User),
			Pool:      common.GeneralizePoolAddress(status.Stratum.Url),
			Uptime:    int(time.Now().Sub(time.Unix(int64(status.StartTime), 0)).Milliseconds() / 1000),
			Worker:    common.ExtractWorkerNameFromAddress(status.Stratum.User),
		},
		Devices: func() []dto.Device {
			devices := make([]dto.Device, 0)
			for _, dev := range status.Miner.Devices {
				devices = append(devices, dto.Device{
					GpuId:            dev.Id,
					Name:             common.GeneralizeGpuName(cropGpuModel(dev.Info)),
					Hashrate:         int(dev.HashrateRaw),
					FanSpeed:         dev.Fan,
					CoreTemp:         dev.Temperature,
					MemoryTemp:       dev.MemTemperature,
					PowerConsumption: dev.Power,
					ShareAccepted:    dev.AcceptedShares,
					ShareRejected:    dev.RejectShares,
					ShareStale:       dev.InvalidShares,
					LhrRate:          common.ValueNotSet,
					CoreClock:        dev.CoreClock,
					MemoryClock:      dev.MemClock,
					CoreUtilization:  dev.CoreUtilization,
					MemUtilization:   dev.MemUtilization,
				})
			}
			return devices
		}(),
	}
	return &commonStatus, nil
}

func cropGpuModel(nbDeviceInfo string) string {
	r, _ := regexp.Compile("(GTX|RTX) \\d+( Ti|)")

	trimmed := strings.TrimSpace(r.FindString(nbDeviceInfo))
	if trimmed != "" {
		return strings.Join(strings.Split(trimmed, " ")[1:], " ")
	}
	return nbDeviceInfo
}
