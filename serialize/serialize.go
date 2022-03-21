package fsrepo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/TRON-US/go-btfs-config"

	"github.com/facebookgo/atomicfile"
)

// ErrNotInitialized is returned when we fail to read the config because the
// repo doesn't exist.
var ErrNotInitialized = errors.New("ipfs not initialized, please run 'ipfs init'")

// ReadConfigFile reads the config from `filename` into `cfg`.
func ReadConfigFile(filename string, cfg interface{}) error {
	f, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			err = ErrNotInitialized
		}
		return err
	}
	defer f.Close()
	if err := json.NewDecoder(f).Decode(cfg); err != nil {
		return fmt.Errorf("failure to decode config: %s", err)
	}
	return nil
}

// WriteConfigFile writes the config from `cfg` into `filename`.
func WriteConfigFile(filename string, cfg interface{}) error {
	err := os.MkdirAll(filepath.Dir(filename), 0755)
	if err != nil {
		return err
	}

	f, err := atomicfile.New(filename, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	return encode(f, cfg)
}

// encode configuration with JSON
func encode(w io.Writer, value interface{}) error {
	// need to prettyprint, hence MarshalIndent, instead of Encoder
	buf, err := config.Marshal(value)
	if err != nil {
		return err
	}

	// rewrite BTFS mounts config with IPFS mounts config
	var cfg config.Config
	if err := json.Unmarshal(buf, &cfg); err != nil {
		return err
	}
	cfg.Mounts.BTFS = cfg.Mounts.IPFS
	cfg.Mounts.BTNS = cfg.Mounts.IPNS
	cfg.Mounts.IPFS = ""
	cfg.Mounts.IPNS = ""
	buf, _ = config.Marshal(cfg)

	_, err = w.Write(buf)
	return err
}

// Load reads given file and returns the read config, or error.
func Load(filename string) (*config.Config, error) {
	var cfg config.Config
	err := ReadConfigFile(filename, &cfg)
	if err != nil {
		return nil, err
	}

	// IPFS and IPNS field value will be used when old config firstly be loaded
	// BTFS and BTNS field value will be used after new config field be written into file
	if cfg.Mounts.BTFS != "" {
		cfg.Mounts.IPFS = cfg.Mounts.BTFS
	}
	if cfg.Mounts.BTNS != "" {
		cfg.Mounts.IPNS = cfg.Mounts.BTNS
	}

	return &cfg, err
}
