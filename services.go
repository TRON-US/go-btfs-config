package config

type Services struct {
	StatusServerDomain string
	OnlineServerDomain string
	HubDomain          string
	EscrowDomain       string
	GuardDomain        string
	ExchangeDomain     string
	SolidityDomain     string
	FullnodeDomain     string
	TrongridDomain     string

	EscrowPubKeys []string
	GuardPubKeys  []string
}
