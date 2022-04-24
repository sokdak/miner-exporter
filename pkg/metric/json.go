package metric

import (
	"fmt"
	"github.com/sokdak/miner-exporter/pkg/dto"
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

	return status, nil
}
