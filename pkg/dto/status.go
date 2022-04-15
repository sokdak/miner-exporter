package dto

type Status struct {
	Miner   Miner    `json:"miner"`
	Devices []Device `json:"devices"`
}
