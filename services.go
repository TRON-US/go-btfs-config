package config

type Services struct {
	StatusServerDomain string
	HubDomain          string
	EscrowDomain       string
	GuardDomain        string
	ExchangeDomain     string
	SolidityDomain     string

	EscrowPubKeys []string
	GuardPubKeys  []string
}
