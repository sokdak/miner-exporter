package smi

const (
	GathererKeyNvidia = "nvidia"
)

type Data struct {
	CoreUtilization   float32
	MemoryUtilization float32
	CoreClock         int
	MemoryClock       int
}
