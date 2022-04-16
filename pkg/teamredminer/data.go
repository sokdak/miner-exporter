package teamredminer

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sokdak/miner-exporter/pkg/common"
	"github.com/sokdak/miner-exporter/pkg/dto"
	"math"
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
	commonStatus := dto.Status{
		Miner: dto.Miner{
			Name:      report.Version.Miner,
			Version:   report.Stat.Generic().MinerVersion,
			Algorithm: "Ethash",
			Address:   report.Pools[0].User,
			Pool:      report.Pools[0].URL,
			Uptime:    int(report.Stat.Generic().Elapsed),
		},
		Devices: func() []dto.Device {
			devs := make([]dto.Device, 0)
			for i, dev := range *report.Devices {
				devs = append(devs, dto.Device{
					GpuId:            report.DeviceDetails[i].Id,
					Name:             common.GeneralizeGpuName(report.DeviceDetails[i].Model),
					Hashrate:         int(math.Floor(dev.MHS30s * 1000 * 1000)),
					FanSpeed:         int(dev.FanPercent),
					CoreTemp:         int(dev.Temperature),
					MemoryTemp:       int(dev.TemperatureMemory),
					PowerConsumption: int(dev.PowerConsumption),
					ShareAccepted:    dev.AcceptedShares,
					ShareRejected:    dev.RejectedShares,
					ShareStale:       int(dev.HardwareErrors),
				})
			}
			return devs
		}(),
	}
	return &commonStatus, nil
}