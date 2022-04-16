package trex

type Summary struct {
	Name          string  `json:"name"`
	Version       string  `json:"version"`
	AcceptedCount int     `json:"accepted_count"`
	Algorithm     string  `json:"algorithm"`
	Difficulty    float64 `json:"difficulty"`
	GpuTotal      int     `json:"gpu_total"`
	ActivePool    Pool    `json:"active_pool"`
	Gpus          []Gpu   `json:"gpus"`
	Hashrate      int     `json:"hashrate"`
	RejectedCount int     `json:"rejected_count"`
	SolvedCount   int     `json:"solved_count"`
	Uptime        int     `json:"uptime"`
}

type Pool struct {
	Ping    int    `json:"ping"`
	Retries int    `json:"retries"`
	Url     string `json:"url"`
	User    string `json:"user"`
	Worker  string `json:"worker"`
}

type Gpu struct {
	DeviceId          int     `json:"device_id"`
	Efficiency        string  `json:"efficiency"`
	FanSpeed          int     `json:"fan_speed"`
	Hashrate          int     `json:"hashrate"`
	Intensity         float32 `json:"intensity"`
	Name              string  `json:"name"`
	Temperature       int     `json:"temperature"`
	MemoryTemperature int     `json:"memory_temperature"`
	Shares            Share   `json:"shares"`
	LhrTune           float32 `json:"lhr_tune"`
	Mclock            int     `json:"mclock"`
	Cclock            int     `json:"cclock"`
	Power             int     `json:"power"`
}

type Share struct {
	AcceptedCount int `json:"accepted_count"`
	InvalidCount  int `json:"invalid_count"`
	RejectedCount int `json:"rejected_count"`
	SolvedCount   int `json:"solved_count"`
}
