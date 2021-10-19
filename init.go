package config

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"time"

	hubpb "github.com/tron-us/go-btfs-common/protos/hub"

	ci "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
)

func Init(out io.Writer, nBitsForKeypair int, keyType string, importKey string, mnemonic string, rmOnUnpin bool) (*Config, error) {
	identity, err := IdentityConfig(out, nBitsForKeypair, keyType, importKey, mnemonic)
	if err != nil {
		return nil, err
	}

	bootstrapPeers, err := DefaultBootstrapPeers()
	if err != nil {
		return nil, err
	}

	datastore := DefaultDatastoreConfig()

	conf := &Config{
		API: API{
			HTTPHeaders: map[string][]string{},
		},

		// setup the node's default addresses.
		// NOTE: two swarm listen addrs, one tcp, one utp.
		Addresses: addressesConfig(),

		Datastore: datastore,
		Bootstrap: BootstrapPeerStrings(bootstrapPeers),
		Identity:  identity,
		Discovery: Discovery{
			MDNS: MDNS{
				Enabled:  true,
				Interval: 10,
			},
		},

		Routing: Routing{
			Type: "dht",
		},

		// setup the node mount points.
		Mounts: Mounts{
			IPFS: "/btfs",
			IPNS: "/btns",
		},

		Ipns: Ipns{
			ResolveCacheSize: 128,
		},

		Gateway: Gateway{
			RootRedirect: "",
			Writable:     false,
			NoFetch:      false,
			PathPrefixes: []string{},
			HTTPHeaders: map[string][]string{
				"Access-Control-Allow-Origin":  []string{"*"},
				"Access-Control-Allow-Methods": []string{"GET"},
				"Access-Control-Allow-Headers": []string{"X-Requested-With", "Range", "User-Agent"},
			},
			APICommands: []string{},
		},
		Services: DefaultServicesConfig(),
		Reprovider: Reprovider{
			Interval: "12h",
			Strategy: "all",
		},
		Swarm: SwarmConfig{
			SwarmKey: DefaultSwarmKey,
			ConnMgr: ConnMgr{
				LowWater:    DefaultConnMgrLowWater,
				HighWater:   DefaultConnMgrHighWater,
				GracePeriod: DefaultConnMgrGracePeriod.String(),
				Type:        "basic",
			},
			EnableAutoRelay: DefaultEnableAutoRelay,
		},
		Experimental: Experiments{
			Libp2pStreamMounting: true, // Enabled for remote api
			StorageClientEnabled: true,
			RemoveOnUnpin:        rmOnUnpin,
			HostsSyncEnabled:     DefaultHostsSyncEnabled,
			HostsSyncFlag:        false,
			HostsSyncMode:        DefaultHostsSyncMode.String(),
		},
	}

	return conf, nil
}

// DefaultHostSyncEnabled is the default value for the periodic hosts sync
// from hub
const DefaultHostsSyncEnabled = false

// DefaultHostsSyncMode is the default value for the hosts sync mode
// from hub
const DefaultHostsSyncMode = hubpb.HostsReq_SCORE
const DefaultHostsSyncModeDev = hubpb.HostsReq_TESTNET

// DefaultConnMgrHighWater is the default value for the connection managers
// 'high water' mark
const DefaultConnMgrHighWater = 900

// DefaultConnMgrLowWater is the default value for the connection managers 'low
// water' mark
const DefaultConnMgrLowWater = 600

// DefaultConnMgrGracePeriod is the default value for the connection managers
// grace period
const DefaultConnMgrGracePeriod = time.Second * 20

// DefaultSwarmKey is the default swarm key for mainnet BTFS
const DefaultSwarmKey = `/key/swarm/psk/1.0.0/
/base16/
64ef95289a6b998c776927ed6e33ca8c9202ee47df90141d09f5ffeeb64b8a66`

// DefaultTestnetSwarmKey is the default swarm key for testnet BTFS
const DefaultTestnetSwarmKey = `/key/swarm/psk/1.0.0/
/base16/
d0566ce7e71d880487a89385296ab8a454967e975955ce0e59bff7991d5539d6`

// DefaultSwarmPort is the default swarm discovery port
const DefaultSwarmPort = 4001

const DefaultEnableAutoRelay = true

func addressesConfig() Addresses {
	return Addresses{
		Swarm: []string{
			fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", DefaultSwarmPort),
			fmt.Sprintf("/ip6/::/tcp/%d", DefaultSwarmPort),
			fmt.Sprintf("/ip4/0.0.0.0/udp/%d/quic", DefaultSwarmPort),
			fmt.Sprintf("/ip6/::/udp/%d/quic", DefaultSwarmPort),
		},
		Announce:   []string{},
		NoAnnounce: []string{},
		API:        Strings{"/ip4/127.0.0.1/tcp/5001"},
		Gateway:    Strings{"/ip4/127.0.0.1/tcp/8080"},
		RemoteAPI:  Strings{"/ip4/127.0.0.1/tcp/5101"},
	}
}

// DefaultDatastoreConfig is an internal function exported to aid in testing.
func DefaultDatastoreConfig() Datastore {
	return Datastore{
		StorageMax:         "10GB",
		StorageGCWatermark: 90, // 90%
		GCPeriod:           "1h",
		BloomFilterSize:    0,
		Spec:               flatfsSpec(),
	}
}

func badgerSpec() map[string]interface{} {
	return map[string]interface{}{
		"type":   "measure",
		"prefix": "badger.datastore",
		"child": map[string]interface{}{
			"type":       "badgerds",
			"path":       "badgerds",
			"syncWrites": false,
			"truncate":   true,
		},
	}
}

func flatfsSpec() map[string]interface{} {
	return map[string]interface{}{
		"type": "mount",
		"mounts": []interface{}{
			map[string]interface{}{
				"mountpoint": "/blocks",
				"type":       "measure",
				"prefix":     "flatfs.datastore",
				"child": map[string]interface{}{
					"type":      "flatfs",
					"path":      "blocks",
					"sync":      true,
					"shardFunc": "/repo/flatfs/shard/v1/next-to-last/2",
				},
			},
			map[string]interface{}{
				"mountpoint": "/",
				"type":       "measure",
				"prefix":     "leveldb.datastore",
				"child": map[string]interface{}{
					"type":        "levelds",
					"path":        "datastore",
					"compression": "none",
				},
			},
		},
	}
}

// DefaultServicesConfig returns the default set of configs for external services.
func DefaultServicesConfig() Services {
	return Services{
		StatusServerDomain: "https://status.btfs.io",
		HubDomain:          "https://hub.btfs.io",
		EscrowDomain:       "https://escrow.btfs.io",
		GuardDomain:        "https://guard.btfs.io",
		ExchangeDomain:     "https://exchange.bt.co",
		SolidityDomain:     "grpc.trongrid.io:50052",
		FullnodeDomain:     "grpc.trongrid.io:50051",
		TrongridDomain:     "https://api.trongrid.io",
		EscrowPubKeys:      []string{"CAISIQPAfB2Mt2ic+n3JcL4vrKXxBCmB0iNh+5BYiXdJNWed/Q=="},
		GuardPubKeys:       []string{"CAISIQJ16EiwvGko4SaBEEUFyMdNZp1vKsTLgIXCY6fRa3/Obg=="},
	}
}

// DefaultServicesConfigDev returns the default set of configs for dev external services.
func DefaultServicesConfigDev() Services {
	return Services{
		StatusServerDomain: "https://status-dev.btfs.io",
		HubDomain:          "https://hub-dev.btfs.io",
		EscrowDomain:       "https://escrow-dev.btfs.io",
		GuardDomain:        "https://guard-dev.btfs.io",
		ExchangeDomain:     "https://exchange-dev.bt.co",
		SolidityDomain:     "grpc.trongrid.io:50052",
		FullnodeDomain:     "grpc.trongrid.io:50051",
		TrongridDomain:     "https://api.shasta.trongrid.io",
		EscrowPubKeys:      []string{"CAISIQJOcRK0q4TOwpswAkvMMq33ksQfhplEyhHcZnEUFbthQg=="},
		GuardPubKeys:       []string{"CAISIQJhPBQWKPPjYcuPWR9sl+QlN0wJSRbQs3yUKmggvubXwg=="},
	}
}

// DefaultServicesConfigTestnet returns the default set of configs for testnet external services.
func DefaultServicesConfigTestnet() Services {
	return Services{
		StatusServerDomain: "https://status-staging.btfs.io",
		HubDomain:          "https://hub-staging.btfs.io",
		EscrowDomain:       "https://escrow-staging.btfs.io",
		GuardDomain:        "https://guard-staging.btfs.io",
		ExchangeDomain:     "https://exchange-staging.bt.co",
		SolidityDomain:     "grpc.trongrid.io:50052",
		FullnodeDomain:     "grpc.trongrid.io:50051",
		TrongridDomain:     "https://api.shasta.trongrid.io",
		EscrowPubKeys:      []string{"CAISIQJOcRK0q4TOwpswAkvMMq33ksQfhplEyhHcZnEUFbthQg=="},
		GuardPubKeys:       []string{"CAISIQJhPBQWKPPjYcuPWR9sl+QlN0wJSRbQs3yUKmggvubXwg=="},
	}
}

// IdentityConfig initializes a new identity.
func IdentityConfig(out io.Writer, nbits int, keyType string, importKey string, mnemonic string) (Identity, error) {
	// TODO guard higher up
	ident := Identity{}

	if nbits < ci.MinRsaKeyBits {
		return ident, ci.ErrRsaKeyTooSmall
	}

	var sk ci.PrivKey
	var pk ci.PubKey
	var err error
	if importKey == "" {
		var key int

		switch keyType {
		case "RSA":
			key = ci.RSA
		case "Ed25519":
			key = ci.Ed25519
		case "Secp256k1":
			key = ci.Secp256k1
		case "ECDSA":
			key = ci.ECDSA
		default:
			key = ci.Secp256k1
			keyType = "Secp256k1"
		}

		fmt.Fprintf(out, "generating %v-bit %s keypair...", nbits, keyType)
		sk, pk, err = ci.GenerateKeyPair(key, nbits)
	} else {
		fmt.Fprintf(out, "generating btfs node keypair with TRON key...")
		skBytes, err := hex.DecodeString(importKey)
		if err != nil {
			return ident, errors.New("cannot decode importKey from a string to byte array")
		}
		sk, err = ci.UnmarshalSecp256k1PrivateKey(skBytes)
		if err != nil {
			return ident, err
		}
		pk = sk.GetPublic()
	}

	if err != nil {
		return ident, err
	}
	fmt.Fprintf(out, "done\n")

	// currently storing key unencrypted. in the future we need to encrypt it.
	// TODO(security)
	skbytes, err := sk.Bytes()
	if err != nil {
		return ident, err
	}
	ident.PrivKey = base64.StdEncoding.EncodeToString(skbytes)
	ident.Mnemonic = mnemonic

	id, err := peer.IDFromPublicKey(pk)
	if err != nil {
		return ident, err
	}
	ident.PeerID = id.Pretty()
	fmt.Fprintf(out, "peer identity: %s\n", ident.PeerID)
	return ident, nil
}
