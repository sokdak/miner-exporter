package teamredminer

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sokdak/miner-exporter/pkg/common"
	"github.com/sokdak/miner-exporter/pkg/dto"
	"math"
	"regexp"
	"strings"
)

func (c Client) Fetch() (interface{}, error) {
	sum, err := c.Client.Summary()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get summary")
	}
	pool, err := c.Client.Pools()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get pools")
	}
	stat, err := c.Client.Stats()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get stats")
	}
	version, err := c.Client.Version()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get version")
	}
	devs, err := c.Client.Devs()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get devs")
	}
	devDetails, err := c.Client.DevDetails()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get devdetails")
	}

	report := &Report{
		Summary:       sum,
		Version:       version,
		Devices:       devs,
		Stat:          stat,
		Pools:         pool,
		DeviceDetails: devDetails,
	}
	return *report, nil
}

func (c Client) Parse(value interface{}) (*dto.Status, error) {
	report, ok := value.(Report)
	if !ok {
		return nil, fmt.Errorf("failed to parse tredminer response dto")
	}

	if common.ExtractWorkerNameFromAddress(report.Pools[0].User) == "" {
		return nil, fmt.Errorf("failed to get worker name")
	}

	commonStatus := dto.Status{
		Miner: dto.Miner{
			Name:      extractName(report.Version.Miner),
			Version:   extractVersionFromName(report.Version.Miner),
			Algorithm: common.GeneralizeAlgorithm("ethash"),
			Address:   common.ExtractAddress(report.Pools[0].User),
			Pool:      common.GeneralizePoolAddress(report.Pools[0].URL),
			Uptime:    int(report.Stat.Generic().Elapsed),
			Worker:    common.ExtractWorkerNameFromAddress(report.Pools[0].User),
		},
		Devices: func() []dto.Device {
			devs := make([]dto.Device, 0)
			for i, dev := range *report.Devices {
				devs = append(devs, dto.Device{
					GpuId:            report.DeviceDetails[i].Id,
					Name:             common.GeneralizeGpuName(cropGpuModel(report.DeviceDetails[i].Model)),
					Hashrate:         int(math.Floor(dev.MHS30s * 1000 * 1000)),
					FanSpeed:         int(dev.FanPercent),
					CoreTemp:         int(dev.Temperature),
					MemoryTemp:       int(dev.TemperatureMemory),
					PowerConsumption: int(dev.PowerConsumption),
					ShareAccepted:    dev.AcceptedShares,
					ShareRejected:    dev.RejectedShares,
					ShareStale:       int(dev.HardwareErrors),
					LhrRate:          common.ValueNotSet,
					CoreClock:        common.GetNonValueInsteadIfNotPresent(int(dev.GPUClock)),
					MemoryClock:      common.GetNonValueInsteadIfNotPresent(int(dev.MemoryClock)),
					CoreUtilization:  common.GetNonValueInsteadIfNotPresent(int(dev.Utility * 100)),
					MemUtilization:   common.ValueNotSet,
				})
			}
			return devs
		}(),
	}
	return &commonStatus, nil
}

func extractName(source string) string {
	return strings.Split(source, " ")[0]
}

func extractVersionFromName(source string) string {
	strs := strings.Split(source, " ")
	if len(strs) < 2 {
		return source
	}
	return strs[1]
}

func cropGpuModel(trmDevModel string) string {
	r, _ := regexp.Compile("(\\[.+\\]|RX \\d+( XT|))")
	if r.MatchString(trmDevModel) {
		return strings.ReplaceAll(strings.ReplaceAll(r.FindString(trmDevModel), "[", ""), "]", "")
	}
	return trmDevModel
}
