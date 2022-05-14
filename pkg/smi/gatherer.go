package smi

type Gatherer interface {
	GetName() string
	Init() error
	Teardown() error

	GetSmiData(deviceIndex int) (*Data, error)
}
