package config

import (
	"errors"
	"fmt"

	peer "github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
)

// DefaultBootstrapAddresses are the hardcoded bootstrap addresses
// for IPFS. they are nodes run by the IPFS team. docs on these later.
// As with all p2p networks, bootstrap is an important security concern.
//
// NOTE: This is here -- and not inside cmd/btfs/init.go -- because of an
// import dependency issue. TODO: move this into a config/default/ package.
var DefaultBootstrapAddresses = []string{
	"/ip4/18.237.54.123/tcp/4001/p2p/QmWJWGxKKaqZUW4xga2BCzT5FBtYDL8Cc5Q5jywd6xPt1g",
	"/ip4/54.213.128.120/tcp/4001/p2p/QmWm3vBCRuZcJMUT9jDZysoYBb66aokmSReX26UaMk8qq5",
	"/ip4/34.213.5.20/tcp/4001/p2p/QmQVQBsM7uoJy8hATjTm51uSAkx2y3iGLhSwA6LWLa7iQJ",
	"/ip4/18.237.202.91/tcp/4001/p2p/QmbVFdiNkvxtc7Nni7yBWAgtHg8MuyhaZ5mDaYR2ZrhhvN",
	"/ip4/13.229.45.41/tcp/4001/p2p/QmX7RZXh27AX8iv2BKLGMgPBiuUpEy8p4LFXgtXAfaZDn9",
	"/ip4/54.254.227.188/tcp/4001/p2p/QmYqCq3PasrzLr3PxtLo5D6spEAJ836W9Re9Eo4zUou45U",
	"/ip4/52.77.240.134/tcp/4001/p2p/QmURPwdLYesWUDB66EGXvDvwcyV44rVRqV2iGNqKN24eVu",
	"/ip4/3.126.224.22/tcp/4001/p2p/QmWTTmvchTodUaVvuKZMo67xk7ZgkxJf4nBo7SZry3vGU5",
	"/ip4/18.194.71.27/tcp/4001/p2p/QmYHkY5CrWcvgaDo4PfvzTQgaZtfaqRGDjwW1MrHUj8cLK",
	"/ip4/54.93.47.134/tcp/4001/p2p/QmeHaHe7WvjeY37z5MYC3qYQcQcuvDwUhwTXtP3KhKLXXK",
}
var DefaultTestnetBootstrapAddresses = []string{
	"/ip4/54.151.185.243/tcp/35369/p2p/16Uiu2HAm8ZYAmsEXE6qBoSeRTeGGR2KVPUZBLBf1dEh9y28ANMMt",
	"/ip4/54.151.185.243/tcp/39789/p2p/16Uiu2HAm56ujXswzAXxY8xT5mz6aPsXUWxwNLYG7bqB2fxFmoAKD",
	"/ip4/54.151.185.243/tcp/34233/p2p/16Uiu2HAm6K1GrRSaCy83N5s9pUbaFp45pHTqhkDvui8gXNMnCPj9",
	"/ip4/54.151.185.243/tcp/39447/p2p/16Uiu2HAmLHxH8qXMMj6pBynFefgFNEzsbaCHuX4hfwaLoRHykjCz",
	"/ip4/18.158.67.141/tcp/41265/p2p/16Uiu2HAmLgKPQUW4kTvrdhnqXtioZZsTeyNxyeXUDnfcFJQxUjKN",
	"/ip4/18.158.67.141/tcp/36007/p2p/16Uiu2HAmJPmu9hUM4NhDYBvkXZidhxf4KA6endccGvKEvujPg3C4",
	"/ip4/18.158.67.141/tcp/39985/p2p/16Uiu2HAm5TuvgzRE56gaZe31J4ErXSvXvxfFhjeYqMybKyR6kyr1",
	"/ip4/18.158.67.141/tcp/35839/p2p/16Uiu2HAkwvgHtaBoKfGDJ9z2B3RvbvMfiBqJX8iS7eR6NFuyS7HF",
	"/ip4/18.163.235.175/tcp/40165/p2p/16Uiu2HAkx4hFxXHxMWF5Sx2SvaV8YgGdzKaZcpbyArixhY9vJcJV",
	"/ip4/18.163.235.175/tcp/34059/p2p/16Uiu2HAmJAxe9pjbiMDRPUFNHbYs8LMfwCaC1sBfWrTnoqnQ38Fh",
	"/ip4/18.163.235.175/tcp/42895/p2p/16Uiu2HAmCuGqf7zHuJMBWtWxUHgYf3WkJqfd7sUHnwsPiBbmJTAj",
	"/ip4/18.163.235.175/tcp/33463/p2p/16Uiu2HAmRNMZmbv3WzWHqaj5CietkmzzzrsvhHK9xnjBjJccerYx",
	"/ip4/18.224.174.215/tcp/39341/p2p/16Uiu2HAm5eXHtcGv6KGZ9QeKM4gbBzxQPr6WLEE6qZpF3AMwzKcF",
	"/ip4/18.224.174.215/tcp/43057/p2p/16Uiu2HAkumR1Ee1HvvHwNqVrCn4jwhDhRNENJVUeXjpTAkx6nNcR",
	"/ip4/18.224.174.215/tcp/44715/p2p/16Uiu2HAm6JRfLXPToTt2hwVcZqTGiG7ZLj25aFbAie7Xeab2H49S",
	"/ip4/18.224.174.215/tcp/33931/p2p/16Uiu2HAmKvjzAPQY5kzjpEJGVXDjMpSC1yNd7ZzLZYBKJs7ydHKV",
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
