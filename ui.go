package config

type UI struct {
	Host   HostUI
	Renter RenterUI
	Wallet WalletUI
}

type HostUI struct {
	Initialized     bool
	ContractManager *ContractManager
}

type ContractManager struct {
	LowWater  int
	HighWater int
	Threshold int64
}

type RenterUI struct {
	Initialized bool
}

type WalletUI struct {
	Initialized bool
}
