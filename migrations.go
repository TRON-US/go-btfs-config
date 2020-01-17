package config

import (
	"reflect"
	"strings"
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
	return updated
}
