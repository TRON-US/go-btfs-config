package config

import (
	"errors"
	"fmt"

	peer "github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"

	// Needs to be imported so that users can import this package directly
	// and still parse the bootstrap addresses.
	_ "github.com/multiformats/go-multiaddr-dns"
)

// DefaultBootstrapAddresses are the hardcoded bootstrap addresses
// for IPFS. they are nodes run by the IPFS team. docs on these later.
// As with all p2p networks, bootstrap is an important security concern.
//
// NOTE: This is here -- and not inside cmd/btfs/init.go -- because of an
// import dependency issue. TODO: move this into a config/default/ package.
var DefaultBootstrapAddresses = []string{
	"/ip4/18.237.54.123/tcp/4001/btfs/QmWJWGxKKaqZUW4xga2BCzT5FBtYDL8Cc5Q5jywd6xPt1g",
	"/ip4/54.213.128.120/tcp/4001/btfs/QmWm3vBCRuZcJMUT9jDZysoYBb66aokmSReX26UaMk8qq5",
	"/ip4/34.213.5.20/tcp/4001/btfs/QmQVQBsM7uoJy8hATjTm51uSAkx2y3iGLhSwA6LWLa7iQJ",
	"/ip4/18.237.202.91/tcp/4001/btfs/QmbVFdiNkvxtc7Nni7yBWAgtHg8MuyhaZ5mDaYR2ZrhhvN",
	"/ip4/13.229.45.41/tcp/4001/btfs/QmX7RZXh27AX8iv2BKLGMgPBiuUpEy8p4LFXgtXAfaZDn9",
	"/ip4/54.254.227.188/tcp/4001/btfs/QmYqCq3PasrzLr3PxtLo5D6spEAJ836W9Re9Eo4zUou45U",
	"/ip4/52.77.240.134/tcp/4001/btfs/QmURPwdLYesWUDB66EGXvDvwcyV44rVRqV2iGNqKN24eVu",
	"/ip4/18.196.49.234/tcp/4001/btfs/QmWTTmvchTodUaVvuKZMo67xk7ZgkxJf4nBo7SZry3vGU5",
	"/ip4/18.194.71.27/tcp/4001/btfs/QmYHkY5CrWcvgaDo4PfvzTQgaZtfaqRGDjwW1MrHUj8cLK",
	"/ip4/54.93.47.134/tcp/4001/btfs/QmeHaHe7WvjeY37z5MYC3qYQcQcuvDwUhwTXtP3KhKLXXK",
}
var DefaultTestnetBootstrapAddresses = []string{
	"/ip4/13.59.69.165/tcp/45301/btfs/16Uiu2HAmFFwNdgSoLhfgJUPEfPEVodppRxaeZBVpAvrH5s3qSkWo",
	"/ip4/13.59.69.165/tcp/34237/btfs/16Uiu2HAmDigS3SDx6g9Sp6MUfdFHvDwS8Zw8E14V6bLhCAHA3jjB",
	"/ip4/13.59.69.165/tcp/43097/btfs/16Uiu2HAm7HQoEbQe1fYt4LtnG6z5TqwTrrqUv5xsnt4nukskWmAi",
	"/ip4/13.59.69.165/tcp/38955/btfs/16Uiu2HAm5WrYvkJwaRP7ZAroWCfjaUxKkNssqcSmEmKJ8vXVYp1o",
	"/ip4/13.59.69.165/tcp/43113/btfs/16Uiu2HAkxwtWLvxtu1Av4ia57qzbuBN2gFJ9zxcaGM6yDtyZxx7T",
	"/ip4/13.229.73.63/tcp/36707/btfs/16Uiu2HAmDis3wAorW46YyNmXNk963VAAHwZ1phjHXj5yduyawAUy",
	"/ip4/13.229.73.63/tcp/42741/btfs/16Uiu2HAmSfqLCyqH5qQQF8zpzPMQvWiQunhWpYtSxwGw5QR2jhgU",
	"/ip4/13.229.73.63/tcp/37403/btfs/16Uiu2HAmBHwyRUETsGqjYpgPRpnMC9y39tcVYH6vKxZidCBcBeFG",
	"/ip4/13.229.73.63/tcp/37739/btfs/16Uiu2HAm2oKy37KvYmiv1nnRWZwUoLPZumNKFxPzhM1t8F3KxADu",
	"/ip4/13.229.73.63/tcp/38869/btfs/16Uiu2HAkz9ayzU2yXBMKuajyFAJ38DKWJhUDMWWehvEHnBYqPwqw",
	"/ip4/3.126.51.74/tcp/40155/btfs/16Uiu2HAmTMEqndByECXuxk1Rg8szxMqwS3tUFFWhAUduFzwfwmfK",
	"/ip4/3.126.51.74/tcp/44569/btfs/16Uiu2HAmL4QNi68nSNbedUWp1A1cRR3z3NuJqQYmAYoj19ht6iNv",
	"/ip4/3.126.51.74/tcp/39703/btfs/16Uiu2HAkzF6JMx4EL2C4cLoCLyQH8t1sgyttQxPfQtNt5FZhvpxs",
	"/ip4/3.126.51.74/tcp/46713/btfs/16Uiu2HAm85HXJA7xmgNxxTVdFRuRCGstvrY8nW6KqfTtkuZrZg64",
	"/ip4/3.126.51.74/tcp/38131/btfs/16Uiu2HAmLvC6gRHir8FCB63seobydTTuwmxMcrBceEiBcNVZBaPw",
}

// ErrInvalidPeerAddr signals an address is not a valid peer address.
var ErrInvalidPeerAddr = errors.New("invalid peer address")

func (c *Config) BootstrapPeers() ([]peer.AddrInfo, error) {
	return ParseBootstrapPeers(c.Bootstrap)
}

// DefaultBootstrapPeers returns the (parsed) set of default bootstrap peers.
// if it fails, it returns a meaningful error for the user.
// This is here (and not inside cmd/btfs/init) because of module dependency problems.
func DefaultBootstrapPeers() ([]peer.AddrInfo, error) {
	ps, err := ParseBootstrapPeers(DefaultBootstrapAddresses)
	if err != nil {
		return nil, fmt.Errorf(`Failed to parse hardcoded bootstrap peers: %s
This is a problem with the BTFS codebase.
Please report it to https://github.com/TRON-US/go-btfs/issues.`, err)
	}
	return ps, nil
}
func DefaultTestnetBootstrapPeers() ([]peer.AddrInfo, error) {
	ps, err := ParseBootstrapPeers(DefaultTestnetBootstrapAddresses)
	if err != nil {
		return nil, fmt.Errorf(`Failed to parse hardcoded testnet bootstrap peers: %s
This is a problem with the BTFS codebase.
Please report it to https://github.com/TRON-US/go-btfs/issues.`, err)
	}
	return ps, nil
}

func (c *Config) SetBootstrapPeers(bps []peer.AddrInfo) {
	c.Bootstrap = BootstrapPeerStrings(bps)
}

// ParseBootstrapPeer parses a bootstrap list into a list of AddrInfos.
func ParseBootstrapPeers(addrs []string) ([]peer.AddrInfo, error) {
	maddrs := make([]ma.Multiaddr, len(addrs))
	for i, addr := range addrs {
		var err error
		maddrs[i], err = ma.NewMultiaddr(addr)
		if err != nil {
			return nil, err
		}
	}
	return peer.AddrInfosFromP2pAddrs(maddrs...)
}

// BootstrapPeerStrings formats a list of AddrInfos as a bootstrap peer list
// suitable for serialization.
func BootstrapPeerStrings(bps []peer.AddrInfo) []string {
	bpss := make([]string, 0, len(bps))
	for _, pi := range bps {
		addrs, err := peer.AddrInfoToP2pAddrs(&pi)
		if err != nil {
			// programmer error.
			panic(err)
		}
		for _, addr := range addrs {
			bpss = append(bpss, addr.String())
		}
	}
	return bpss
}
