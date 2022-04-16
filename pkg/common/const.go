package common

import "time"

const (
	MinerTypeGMiner       = "gminer"
	MinerTypeTeamRedMiner = "teamredminer"
	MinerTypeTrexMiner    = "trex"

	InitRetryBackoff   = 10 * time.Second
	InitRetryThreshold = 60
)
