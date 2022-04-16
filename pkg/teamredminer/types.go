package teamredminer

import "github.com/sokdak/go-teamredminer-api"

type Report struct {
	Summary       *cgminer.Summary
	Version       *cgminer.Version
	Devices       *[]cgminer.Devs
	Pools         []cgminer.Pool
	Stat          cgminer.Stats
	DeviceDetails []cgminer.DeviceDetail
}
