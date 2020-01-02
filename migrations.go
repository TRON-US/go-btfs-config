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

func migrate_3_StorageSettings(cfg *Config, fromV0 bool) bool {
	// 1) Enable host if user opted in (analytics = true) AND
	// it is a new upgrade from 0.x.x version
	// 2) Enable renter if it is a new upgrade from 0.x.x version
	if !fromV0 {
		return false
	}
	// no error possible
	Profiles["storage-client"].Transform(cfg)
	if cfg.Experimental.Analytics {
		Profiles["storage-host"].Transform(cfg)
	}
	return true
}

// MigrateConfig migrates config options to the latest known version
// It may correct incompatible configs as well
func MigrateConfig(cfg *Config) bool {
	updated := false
	upToV1 := migrate_1_Services(cfg)
	updated = upToV1 || updated
	updated = migrate_2_StatusUrl(cfg) || updated
	updated = migrate_3_StorageSettings(cfg, upToV1) || updated
	return updated
}
