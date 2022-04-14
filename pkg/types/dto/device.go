package dto

type Device struct {
	GpuId            int    `json:"gpu_id"`
	Name             string `json:"name"`
	Hashrate         uint64 `json:"hashrate"`
	CoreTemp         uint32 `json:"core_temp"`
	MemoryTemp       uint32 `json:"memory_temp"`
	PowerConsumption uint32 `json:"power_consumption"`
	ShareAccepted    uint32 `json:"share_accepted"`
	ShareRejected    uint32 `json:"share_rejected"`
	ShareStale       uint32 `json:"share_stale"`
}
