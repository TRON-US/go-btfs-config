package config

import (
	"reflect"
	"strings"

	"github.com/libp2p/go-libp2p-core/peer"
)

func migrate_1_Services(cfg *Config) bool {
	emptyServices := Services{}
	if reflect.DeepEqual(cfg.Services, emptyServices) {
		cfg.Services = DefaultServicesConfig()
		return true
	}
	if len(cfg.Services.EscrowPubKeys) == 0 || len(cfg.Services.GuardPubKeys) == 0 {
		cfg.Services = DefaultServicesConfig()
		return true
	}
	return false
}

func migrate_2_StatusUrl(cfg *Config) bool {
	if strings.Contains(cfg.Services.StatusServerDomain, "db.btfs.io") {
		ds := DefaultServicesConfig()
		cfg.Services.StatusServerDomain = ds.StatusServerDomain
		return true
	}
	return false
}

func migrate_3_StorageSettings(cfg *Config, fromV0, inited, hasHval bool) bool {
	// 1) Enable host
	//    a) Upgrade from 0.x.x -> 1.x.x and has hval (bt client)
	//    b) New profile and has hval (bt client)
	// 2) Enable renter if it is a new upgrade from 0.x.x version
	if fromV0 {
		Profiles["storage-client"].Transform(cfg)
	}
	if hasHval && (fromV0 || inited) {
		Profiles["storage-host"].Transform(cfg)
	}
	return true
}

func migrate_4_SwarmKey(cfg *Config) bool {
	if cfg.Swarm.SwarmKey == "" {
		cfg.Swarm.SwarmKey = DefaultSwarmKey
		return true
	}
	return false
}

// checks to see if the current config contains known obsolete ip addresses
// for bootstrap nodes.
// Replaces all bootstrap nodes with default values if so.
func migrate_5_Bootstrap_node(cfg *Config) bool {
	// Only migrate on prod settings
	if cfg.Swarm.SwarmKey != DefaultSwarmKey {
		return false
	}
	obns := []string{
		"3.120.224.94",
		"18.196.49.234",
		"/btfs/", // migrate to ipfs 0.5.0+ protocol
	}
	peers, _ := DefaultBootstrapPeers()
	return doMigrateNodes(cfg, obns, peers)
}

func migrate_6_EnableAutoRelay(cfg *Config) bool {
	if cfg.Swarm.EnableAutoRelay != DefaultEnableAutoRelay {
		cfg.Swarm.EnableAutoRelay = DefaultEnableAutoRelay
		return true
	}
	return false
}

// checks to see if the current config contains known obsolete ip addresses
// for testnet bootstrap nodes.
// Replaces all testnet bootstrap nodes with default values if so.
func migrate_7_Testnet_Bootstrap_node(cfg *Config) bool {
	// Only migrate on testnet settings
	if cfg.Swarm.SwarmKey != DefaultTestnetSwarmKey {
		return false
	}
	obns := []string{
		"52.57.56.230",
		"13.59.69.165/tcp/43113",
		"13.229.73.63/tcp/38869",
		"3.126.51.74/tcp/38131",
		"/btfs/", // migrate to ipfs 0.5.0+ protocol
	}
	peers, _ := DefaultTestnetBootstrapPeers()
	return doMigrateNodes(cfg, obns, peers)
}

func doMigrateNodes(cfg *Config, obsoleteBootstrapNodes []string, defaultPeers []peer.AddrInfo) bool {
	currentBootstrapNodeList := cfg.Bootstrap

	for _, obsoleteNode := range obsoleteBootstrapNodes {
		for _, bootstrapNode := range currentBootstrapNodeList {
			if strings.Contains(bootstrapNode, obsoleteNode) {
				cfg.SetBootstrapPeers(defaultPeers)
				return true
			}
		}
	}
	return false
}

func migrate_8_AnnounceDefault(cfg *Config, beforeV1B2 bool) bool {
	if beforeV1B2 {
		cfg.Addresses.Announce = []string{}
		return true
	}
	return false
}

func migrate_9_WalletDomain(cfg *Config) bool {
	if strings.Contains(cfg.Services.EscrowDomain, "dev") || strings.Contains(cfg.Services.EscrowDomain, "staging") {
		if len(cfg.Services.ExchangeDomain) == 0 {
			ds := DefaultServicesConfigTestnet()
			cfg.Services.ExchangeDomain = ds.ExchangeDomain
			cfg.Services.SolidityDomain = ds.SolidityDomain
			return true
		}
	} else {
		if len(cfg.Services.ExchangeDomain) == 0 {
			ds := DefaultServicesConfig()
			cfg.Services.ExchangeDomain = ds.ExchangeDomain
			cfg.Services.SolidityDomain = ds.SolidityDomain
			return true
		}
	}
	return false
}

// check if HTTPHeaders, API and Gateway config are fully configured, then clean HTTPHeaders
func migrate_10_CleanAPIHTTPHeaders(cfg *Config) bool {
	condCount := 3
	httpHeaderFullConfig := map[string][]string{
		"Access-Control-Allow-Origin":      {"*"},
		"Access-Control-Allow-Methods":     {"PUT", "GET", "POST", "OPTIONS"},
		"Access-Control-Allow-Credentials": {"true"},
	}
	if reflect.DeepEqual(cfg.API.HTTPHeaders, httpHeaderFullConfig) {
		condCount--
	}
	addressesAPIFullConfig := Strings{"/ip4/0.0.0.0/tcp/5001"}
	if reflect.DeepEqual(cfg.Addresses.API, addressesAPIFullConfig) {
		condCount--
	}
	addressesGatewayFullConfig := Strings{"/ip4/0.0.0.0/tcp/8080"}
	if reflect.DeepEqual(cfg.Addresses.Gateway, addressesGatewayFullConfig) {
		condCount--
	}
	if condCount == 0 {
		// clean httpheaders configuration
		cfg.API.HTTPHeaders = make(map[string][]string)
		return true
	}
	return false
}

func migrate_11_ExchangeDomain(cfg *Config) bool {
	// migrate staging domain -> staging
	if strings.Contains(cfg.Services.EscrowDomain, "staging") &&
		strings.Contains(cfg.Services.ExchangeDomain, "dev") {
		ds := DefaultServicesConfigTestnet()
		cfg.Services.ExchangeDomain = ds.ExchangeDomain
		return true
	}
	return false
}

func migrate_12_FullnodeDomain(cfg *Config) bool {
	if strings.Contains(cfg.Services.EscrowDomain, "dev") || strings.Contains(cfg.Services.EscrowDomain, "staging") {
		if len(cfg.Services.FullnodeDomain) == 0 {
			ds := DefaultServicesConfigTestnet()
			cfg.Services.FullnodeDomain = ds.FullnodeDomain
			return true
		}
	} else {
		if len(cfg.Services.FullnodeDomain) == 0 {
			ds := DefaultServicesConfig()
			cfg.Services.FullnodeDomain = ds.FullnodeDomain
			return true
		}
	}
	return false
}

// MigrateConfig migrates config options to the latest known version
// It may correct incompatible configs as well
// inited = just initialized in the same call
// hasHval = passed in Hval in the same call
func MigrateConfig(cfg *Config, inited, hasHval bool) bool {
	updated := false
	upToV1 := migrate_1_Services(cfg)
	updated = upToV1 || updated
	updated = migrate_2_StatusUrl(cfg) || updated
	updated = migrate_3_StorageSettings(cfg, upToV1, inited, hasHval) || updated
	upToV1B2 := migrate_4_SwarmKey(cfg)
	updated = upToV1B2 || updated
	updated = migrate_5_Bootstrap_node(cfg) || updated
	updated = migrate_6_EnableAutoRelay(cfg) || updated
	updated = migrate_7_Testnet_Bootstrap_node(cfg) || updated
	updated = migrate_8_AnnounceDefault(cfg, upToV1B2) || updated
	updated = migrate_9_WalletDomain(cfg) || updated
	updated = migrate_10_CleanAPIHTTPHeaders(cfg) || updated
	updated = migrate_11_ExchangeDomain(cfg) || updated
	updated = migrate_12_FullnodeDomain(cfg) || updated
	return updated
}
