package config

type Experiments struct {
	FilestoreEnabled     bool
	UrlstoreEnabled      bool
	ShardingEnabled      bool
	GraphsyncEnabled     bool
	Libp2pStreamMounting bool
	P2pHttpProxy         bool
	StrategicProviding   bool
	StorageHostEnabled   bool
	StorageClientEnabled bool
	Analytics            bool
	RemoveOnUnpin        bool
	HostsSyncEnabled     bool
	HostsSyncMode        string
	DisableAutoUpdate    bool
	HostRepairEnabled    bool
	HostChallengeEnabled bool
}
