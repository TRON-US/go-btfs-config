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

// MigrateConfig migrates config options to the latest known version
// It may correct incompatible configs as well
func MigrateConfig(cfg *Config) bool {
	updated := false
	updated = migrate_1_Services(cfg) || updated
	updated = migrate_2_StatusUrl(cfg) || updated
	return updated
}
