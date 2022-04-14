package types

type MinerInterface interface {
	NewClient(protocol string, host string, port int)
	Fetch() []byte
	Parse(value string)
}
