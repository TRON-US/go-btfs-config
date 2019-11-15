package config

import (
	"reflect"
)

// MigrateConfig migrates config options to the latest known version
// It may correct incompatible configs as well
func MigrateConfig(cfg *Config) bool {
	emptyServices := Services{}
	if reflect.DeepEqual(cfg.Services, emptyServices) {
		cfg.Services = DefaultServicesConfig()
		return true
	}
	return false
}
