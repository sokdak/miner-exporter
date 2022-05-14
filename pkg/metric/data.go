package metric

import (
	"fmt"
	"github.com/sokdak/miner-exporter/pkg/common"
	"github.com/sokdak/miner-exporter/pkg/dto"
	"github.com/sokdak/miner-exporter/pkg/smi"
)

type StatusOverrideStrategy string

const (
	IfNotPresent  StatusOverrideStrategy = "empty"
	All           StatusOverrideStrategy = "all"
	DoNotOverride StatusOverrideStrategy = ""
)

var (
	UserSetOverrideStrategy = IfNotPresent
)

func fetchAndGetStatus() (*dto.Status, error) {
	m := GetMiner()
	if m == nil {
		return nil, fmt.Errorf("failed to get miner")
	}

	// fetch
	dataFetched, err := m.Fetch()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data")
	}

	// parse
	status, err := m.Parse(dataFetched)
	if err != nil {
		return nil, fmt.Errorf("failed to parse data")
	}

	// inject smi information (inject missing values)
	injectGathererInfo(status, UserSetOverrideStrategy)

	return status, nil
}

func injectGathererInfo(data *dto.Status, strategy StatusOverrideStrategy) *dto.Status {
	var nData *dto.Status
	err := common.DeepCopy(data, nData)
	if err != nil {
		return data
	}

	if strategy != DoNotOverride {
		for idx, dev := range nData.Devices {
			var smiData *smi.Data

			// get smi info
			// TODO: how could distinguish chip manufacturer such as nvidia or amd..?
			// in many cases, all nvidia or all amd rigs are in used. querying not existing index may leads to error
			// lets query all gpu indices to all gatherer and pick not errored one!
			smiData, err = smi.GlobalGathererManager.Gatherers[smi.GathererKeyNvidia].GetSmiData(dev.GpuId)
			if err != nil {
				continue
			}

			// set value by StatusOverrideStrategy
			switch strategy {
			case IfNotPresent:
				// check one-by-one and set value
				if nData.Devices[idx].CoreUtilization == common.ValueNotSet {
					nData.Devices[idx].CoreUtilization = smiData.CoreUtilization
				}
				if nData.Devices[idx].MemUtilization == common.ValueNotSet {
					nData.Devices[idx].MemUtilization = smiData.MemoryUtilization
				}
				if nData.Devices[idx].CoreClock == common.ValueNotSet {
					nData.Devices[idx].CoreClock = smiData.CoreClock
				}
				if nData.Devices[idx].MemoryClock == common.ValueNotSet {
					nData.Devices[idx].MemoryClock = smiData.MemoryClock
				}
			case All:
				// override all values
				nData.Devices[idx].CoreUtilization = smiData.CoreUtilization
				nData.Devices[idx].MemUtilization = smiData.MemoryUtilization
				nData.Devices[idx].CoreClock = smiData.CoreClock
				nData.Devices[idx].MemoryClock = smiData.MemoryClock
			case DoNotOverride:
				// this case was proceeded before getting smi info :)
			}
		}
	}

	return nData
}
