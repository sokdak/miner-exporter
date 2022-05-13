package common

import "time"

const (
	MinerTypeGMiner       = "gminer"
	MinerTypeTeamRedMiner = "teamredminer"
	MinerTypeTrexMiner    = "trex"
	MinerTypeNbMiner      = "nbminer"

	InitRetryBackoff   = 30 * time.Second
	InitRetryThreshold = 120
)

const (
	ValueNotSet = -1
)
