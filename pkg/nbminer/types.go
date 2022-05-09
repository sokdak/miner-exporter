package nbminer

type Status struct {
	Miner       Miner   `json:"miner"`
	RebootTimes int     `json:"reboot_times"`
	StartTime   int     `json:"start_time"`
	Stratum     Stratum `json:"stratum"`
	Version     string  `json:"version"`
}

type Stratum struct {
	Algorithm  string `json:"algorithm"`
	Difficulty string `json:"difficulty"`
	Latency    int    `json:"latency"`
	Url        string `json:"url"`
	User       string `json:"user"`
}

type Device struct {
	CoreClock       int     `json:"core_clock"`
	CoreUtilization int     `json:"core_utilization"`
	Fan             int     `json:"fan"`
	HashrateRaw     float64 `json:"hashrate_raw"`
	Id              int     `json:"id"`
	Info            string  `json:"info"`
	Temperature     int     `json:"temperature"`
	MemTemperature  int     `json:"memTemperature"`
	MemClock        int     `json:"mem_clock"`
	MemUtilization  int     `json:"mem_utilization"`
	PciBusId        int     `json:"pci_bus_id"`
	Power           int     `json:"power"`
	AcceptedShares  int     `json:"accepted_shares"`
	InvalidShares   int     `json:"invalid_shares"`
	RejectShares    int     `json:"reject_shares"`
}

type Miner struct {
	Devices []Device `json:"devices"`
}
