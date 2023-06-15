package config

var (
	// chain ID
	ethChainID      = int64(5)
	tronChainID     = int64(100)
	bttcChainID     = int64(199)
	bttcTestChainID = int64(1029)
	testChainID     = int64(1337)
)

// the configuration of the local node's ChainInfo.
type ChainInfo struct {
	ChainId            int64  `json:",omitempty"`
	CurrentFactory     string `json:",omitempty"`
	PriceOracleAddress string `json:",omitempty"`
	VaultLogicAddress  string `json:",omitempty"`
	Endpoint           string `json:",omitempty"`
}
