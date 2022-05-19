package gminer

type Device struct {
	GpuId             int    `json:"gpu_id"`
	Name              string `json:"name"`
	Speed             int    `json:"speed"`
	AcceptedShares    int    `json:"accepted_shares"`
	RejectedShares    int    `json:"rejected_shares"`
	InvalidShares     int    `json:"invalid_shares"`
	StaleShares       int    `json:"stale_shares"`
	Fan               int    `json:"fan"`
	Temperature       int    `json:"temperature"`
	MemoryTemperature int    `json:"memory_temperature"`
	CoreClock         int    `json:"core_clock,omitempty"`
	MemoryClock       int    `json:"memory_clock,omitempty"`
	PowerUsage        int    `json:"power_usage"`
}

type Stat struct {
	Miner     string   `json:"miner"`
	Uptime    int      `json:"uptime"`
	Server    string   `json:"server"`
	User      string   `json:"user"`
	Algorithm string   `json:"algorithm"`
	Devices   []Device `json:"devices"`
}
