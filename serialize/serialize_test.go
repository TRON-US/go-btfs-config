package fsrepo

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"testing"

	config "github.com/TRON-US/go-btfs-config"
)

func TestConfig(t *testing.T) {
	const filename = ".ipfsconfig"
	cfgWritten := new(config.Config)
	cfgWritten.Identity.PeerID = "faketest"

	err := WriteConfigFile(filename, cfgWritten)
	if err != nil {
		t.Fatal(err)
	}
	cfgRead, err := Load(filename)
	if err != nil {
		t.Fatal(err)
	}
	if cfgWritten.Identity.PeerID != cfgRead.Identity.PeerID {
		t.Fatal()
	}
	st, err := os.Stat(filename)
	if err != nil {
		t.Fatalf("cannot stat config file: %v", err)
	}

	if runtime.GOOS != "windows" { // see https://golang.org/src/os/types_windows.go
		if g := st.Mode().Perm(); g&0117 != 0 {
			t.Fatalf("config file should not be executable or accessible to world: %v", g)
		}
	}
}

func TestMountsConfigPatch(t *testing.T) {
	const (
		newConfigFile     = "./config_new.json"
		changedConfigFile = "./config_changed.json"
	)

	// load old config
	cfgFile0, err := config.Filename(os.Getenv("BTFS_PATH"))
	if err != nil {
		t.Fatalf("get config filename: %v\n", err)
	}
	cfg0, err := Load(cfgFile0)
	if err != nil {
		t.Fatalf("load old config: %v\n", err)
	}
	if cfg0.Mounts.IPFS == "" || cfg0.Mounts.IPNS == "" {
		t.Error("load old config: old mounts fields content is empty")
	}
	t.Logf("old mounts: %+v\n", cfg0.Mounts)

	// write new config
	if err := WriteConfigFile(newConfigFile, cfg0); err != nil {
		t.Fatalf("write new config: %v\n", err)
	}
	read, err := ioutil.ReadFile(newConfigFile)
	if err != nil {
		t.Fatalf("read new config new file: %v\n", err)
	}
	readStr := string(read)
	ctBTFS := fmt.Sprintf(`"BTFS": "%s"`, cfg0.Mounts.IPFS)
	ctBTNS := fmt.Sprintf(`"BTNS": "%s"`, cfg0.Mounts.IPNS)
	if !strings.Contains(readStr, ctBTFS) || !strings.Contains(readStr, ctBTNS) {
		t.Fatal("write new config: new mounts fields not exist or content not match")
	}
	if strings.Contains(readStr, `IPFS: "`) || strings.Contains(readStr, `IPNS: "`) {
		t.Fatal("write new config: old mounts fields still exist")
	}

	// load new config
	cfg1, err := Load(newConfigFile)
	if err != nil {
		t.Fatalf("load new config: %v\n", err)
	}
	if cfg1.Mounts.IPFS != cfg0.Mounts.IPFS || cfg1.Mounts.IPNS != cfg0.Mounts.IPNS {
		t.Fatal("load new config: new mounts fields content not matched with old mounts")
	}
	t.Logf("new mounts: %+v\n", cfg1.Mounts)

	// change config
	const (
		changedBTFSContent = "/changed_btfs"
		changedBTNSContent = "/changed_btns"
	)
	cfg1.Mounts.IPFS = changedBTFSContent
	cfg1.Mounts.IPNS = changedBTNSContent
	if err := WriteConfigFile(changedConfigFile, cfg1); err != nil {
		t.Fatalf("write changed config: %v\n", err)
	}

	// load changed config
	cfg2, err := Load(changedConfigFile)
	if err != nil {
		t.Fatalf("load changed config: %v\n", err)
	}
	if cfg2.Mounts.IPFS != changedBTFSContent || cfg2.Mounts.IPNS != changedBTNSContent {
		t.Fatal("load changed config: changed mounts fields content not matched with changed content")
	}
	t.Logf("changed mounts: %+v", cfg2.Mounts)

	// clear config files
	_ = os.Remove(newConfigFile)
	_ = os.Remove(changedConfigFile)
}
