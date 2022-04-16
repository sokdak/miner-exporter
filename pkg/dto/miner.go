package dto

type Miner struct {
	Name      string `json:"name"`
	Version   string `json:"version,omitempty"`
	Algorithm string `json:"algorithm"`
	Address   string `json:"address"`
	Pool      string `json:"pool"`
	Uptime    int    `json:"uptime"`
	Worker    string `json:"worker"`
}
