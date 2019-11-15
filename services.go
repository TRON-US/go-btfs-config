package config

type Services struct {
	StatusServerDomain string
	HubDomain          string
	EscrowDomain       string
	GuardDomain        string

	EscrowPubKeys []string
	GuardPubKeys  []string
}
