package common

import "time"

const (
	MinerTypeGMiner       = "gminer"
	MinerTypeTeamRedMiner = "teamredminer"
	MinerTypeTrexMiner    = "trex"

	InitRetryBackoff   = 30 * time.Second
	InitRetryThreshold = 120
)
