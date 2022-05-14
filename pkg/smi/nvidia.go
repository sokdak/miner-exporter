package smi

import (
	"github.com/pkg/errors"

	"github.com/NVIDIA/gpu-monitoring-tools/bindings/go/nvml"
	"github.com/go-logr/logr"
)

type Nvml struct {
	Log         logr.Logger
	Initialized bool
}

func NewSmi(log logr.Logger) *Nvml {
	return &Nvml{
		Log: log.WithName("nvidia-smi"),
	}
}

func (g *Nvml) Init() error {
	err := nvml.Init()
	if err != nil {
		return errors.Wrapf(err, "failed to init nvml")
	}

	g.Initialized = true
	return nil
}

func (g *Nvml) Teardown() error {
	err := nvml.Shutdown()
	if err != nil {
		return errors.Wrapf(err, "failed to teardown nvml")
	}
	g.Initialized = false
	return nil
}

func (g *Nvml) GetName() string {
	return "nvidia-smi"
}

func (g *Nvml) GetSmiData(deviceIndex int) (*Data, error) {
	dev, err := nvml.NewDeviceLite(uint(deviceIndex))
	if err != nil {
		g.Log.Error(err, "failed to get device(lite)", "deviceIndex", deviceIndex)
		return nil, err
	}

	status, err := dev.Status()
	if err != nil {
		g.Log.Error(err, "failed to get status", "deviceIndex", deviceIndex)
	}

	return &Data{
		CoreUtilization:   float32(*status.Utilization.GPU),
		MemoryUtilization: float32(*status.Utilization.Memory),
		CoreClock:         int(*status.Clocks.Cores),
		MemoryClock:       int(*status.Clocks.Memory),
	}, nil
}
