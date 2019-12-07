package config

type Experiments struct {
	FilestoreEnabled     bool
	UrlstoreEnabled      bool
	ShardingEnabled      bool
	Libp2pStreamMounting bool
	P2pHttpProxy         bool
	QUIC                 bool
	PreferTLS            bool
	StrategicProviding   bool
	StorageHostEnabled   bool
	StoragePrice         float64
	StorageClientEnabled bool
	Analytics            bool
	RemoveOnUnpin        bool
	HostsSyncEnabled     bool
	HostsSyncMode        string
}
