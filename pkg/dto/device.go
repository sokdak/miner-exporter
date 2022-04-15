package dto

type Device struct {
	GpuId            int    `json:"gpu_id"`
	Name             string `json:"name"`
	Hashrate         int    `json:"hashrate"`
	FanSpeed         int    `json:"fan_speed"`
	CoreTemp         int    `json:"core_temp"`
	MemoryTemp       int    `json:"memory_temp"`
	PowerConsumption int    `json:"power_consumption"`
	ShareAccepted    int    `json:"share_accepted"`
	ShareRejected    int    `json:"share_rejected"`
	ShareStale       int    `json:"share_stale"`
}
