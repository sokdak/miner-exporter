package dto

type Miner struct {
	Name      string `json:"name"`
	Version   string `json:"version"`
	Algorithm string `json:"algorithm"`
	Address   string `json:"address"`
	Pool      string `json:"pool"`
}
