package smi

import (
	"fmt"
	"github.com/go-logr/logr"
)

type GathererManager struct {
	Log       logr.Logger
	Gatherers map[string]Gatherer
}

var (
	GlobalGathererManager = &GathererManager{}
)

func NewGathererManagerAndInit(log logr.Logger) {
	log = log.WithName("device-gatherer")
	gatherers := map[string]Gatherer{}

	// init nvidia smi
	nvml := NewSmi(log)
	if err := nvml.Init(); err != nil {
		log.Error(err, "failed to init device info gatherer", "gatherer", "nvidia-smi")
	}
	gatherers[GathererKeyNvidia] = nvml

	// TODO: implement amd device info (rocm?)

	log.Info(fmt.Sprintf("%d gatherer(s) initialized", len(gatherers)))
	GlobalGathererManager = &GathererManager{
		Log:       log,
		Gatherers: gatherers,
	}
}

func (g *GathererManager) Teardown() {
	for _, gatherer := range g.Gatherers {
		err := gatherer.Teardown()
		if err != nil {
			g.Log.Error(err, "failed to teardown device info gatherer", "gatherer", gatherer.GetName())
		}
	}
}
