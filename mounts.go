package config

// Mounts stores the (string) mount points
type Mounts struct {
	IPFS           string `json:"IPFS,omitempty"`
	IPNS           string `json:"IPNS,omitempty"`
	FuseAllowOther bool

	BTFS string
	BTNS string
}
