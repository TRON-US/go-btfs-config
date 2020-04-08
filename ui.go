package config

type UI struct {
	Host   HostUI
	Renter RenterUI
}

type HostUI struct {
	Initialized bool
}

type RenterUI struct {
	Initialized bool
}
