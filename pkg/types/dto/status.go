package dto

type Status struct {
	TotalHashrate int64    `json:"total_hashrate"`
	Healthy       bool     `json:"healthy"`
	Miner         Miner    `json:"miner"`
	Devices       []Device `json:"devices"`
}
