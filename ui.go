package config

type UI struct {
	Host   HostUI
	Renter RenterUI
	Wallet WalletUI
}

type HostUI struct {
	Initialized bool
}

type RenterUI struct {
	Initialized bool
}

type WalletUI struct {
	Initialized bool
}
