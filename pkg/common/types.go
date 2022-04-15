package common

import (
	"github.com/sokdak/miner-exporter/pkg/dto"
)

type MinerInterface interface {
	Init() error
	Ping() error
	Fetch() (interface{}, error)
	Parse(value interface{}) (*dto.Status, error)
}

type ConnectionInfo struct {
	Protocol string
	Host     string
	Port     int
}
