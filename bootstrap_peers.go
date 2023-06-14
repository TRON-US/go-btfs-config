package config

import (
	"errors"
	"fmt"

	peer "github.com/libp2p/go-libp2p/core/peer"
	ma "github.com/multiformats/go-multiaddr"
)

// DefaultBootstrapAddresses are the hardcoded bootstrap addresses
// for IPFS. they are nodes run by the IPFS team. docs on these later.
// As with all p2p networks, bootstrap is an important security concern.
//
// NOTE: This is here -- and not inside cmd/btfs/init.go -- because of an
// import dependency issue. TODO: move this into a config/default/ package.
var DefaultBootstrapAddresses = []string{
	"/ip4/54.255.27.251/tcp/4001/p2p/QmURPwdLYesWUDB66EGXvDvwcyV44rVRqV2iGNqKN24eVu",
	"/ip4/13.213.254.73/tcp/4001/p2p/QmX7RZXh27AX8iv2BKLGMgPBiuUpEy8p4LFXgtXAfaZDn9",
	"/ip4/52.221.82.136/tcp/4001/p2p/QmYqCq3PasrzLr3PxtLo5D6spEAJ836W9Re9Eo4zUou45U",
	"/ip4/3.1.76.240/tcp/4001/p2p/16Uiu2HAmQfh6CYSWG1MM1DnzJ9duM8jxqn6vxNbjGBsmzc3kkctp",
	"/ip4/34.213.53.108/tcp/4001/p2p/QmWm3vBCRuZcJMUT9jDZysoYBb66aokmSReX26UaMk8qq5",
	"/ip4/35.161.203.220/tcp/4001/p2p/QmWJWGxKKaqZUW4xga2BCzT5FBtYDL8Cc5Q5jywd6xPt1g",
	"/ip4/35.84.213.199/tcp/4001/p2p/QmbVFdiNkvxtc7Nni7yBWAgtHg8MuyhaZ5mDaYR2ZrhhvN",
	"/ip4/35.84.151.38/tcp/4001/p2p/QmQVQBsM7uoJy8hATjTm51uSAkx2y3iGLhSwA6LWLa7iQJ",
	"/ip4/3.66.57.90/tcp/4001/p2p/16Uiu2HAmFT6NXQkzDZXHxuFC4qFt6D1ALf57AkJV9U54HoafX7FX",
	"/ip4/3.69.104.217/tcp/4001/p2p/16Uiu2HAm2v2NBTLYmzVnLJoNbLCmdn29Gv8qLixHCJNeE81rUUYe",
	"/ip4/3.126.224.22/tcp/4001/p2p/16Uiu2HAmNngtNogFpcAUdc6wdSDmb8ZZQjjoDDWaatBXW1rHsYpu",
	"/ip4/35.158.193.90/tcp/4001/p2p/16Uiu2HAmLY4kyhMuoBntyXSt2YssZCHjefHEAXrYVc6acB7KEBh3",
	"/ip4/13.232.5.9/tcp/4001/p2p/16Uiu2HAmSkFDwHU3snrYD2ib5wWeKcsuFMZWEPt31z5YVJ8ktw1p",
	"/ip4/3.7.220.224/tcp/4001/p2p/16Uiu2HAmRVtFaXksAqb8W4Fyr8g5jkggeGDFdVcp8dQ724NMvpcR",
	"/ip4/3.109.125.91/tcp/4001/p2p/16Uiu2HAkzggX1jKwc1xen5qNPQ5RKNkXQqmH2PYAKGd8JZ15YQmK",
	"/ip4/65.1.217.86/tcp/4001/p2p/16Uiu2HAmGwDdvK4jAi1Ahga3zkiuW6HFZKKFNWtqVUFXvaSCNjdg",
	"/ip4/15.184.66.135/tcp/4001/p2p/16Uiu2HAmBQcXzrgo9MVD8xZwt4CrLzPRK1yKAVM7eY9GhXMJYHmC",
	"/ip4/15.184.174.48/tcp/4001/p2p/16Uiu2HAm6Bkxj81JQxa67aja7UWznjTgzAAVzPAqZMVD6oGpw7ST",
	"/ip4/15.184.108.65/tcp/4001/p2p/16Uiu2HAkzRyGYEba2B3SBXdwp328LNFhRG4qhJVZrN6tsJK5KKu5",
	"/ip4/15.184.96.102/tcp/4001/p2p/16Uiu2HAmFVWTvouWpQTRjMb4bUaidfLzsH2RVogcGHb6RwvPSxuT",
	"/ip4/16.24.14.84/tcp/4001/p2p/16Uiu2HAmM96uUH53Ab9JBWfuwUBXJvGMbfVbsBXiGZGqStP93DTS",
	"/ip4/16.24.16.4/tcp/4001/p2p/16Uiu2HAmJ6vEtzmmC6nM6SJwHA9NCPwTRWy7K5WT2UFXDqzJFGSf",
	"/ip4/3.76.64.148/tcp/4001/p2p/16Uiu2HAmFc3snGkwK76yMYMAkHWhq6GD29w7m8Sa7kUciUK5xovu",
	"/ip4/3.78.178.244/tcp/4001/p2p/16Uiu2HAmHeUHakzYG1YWfWoSriVwKhSHYz88rL3USmgeRpqtWqMw",
	"/ip4/3.7.21.138/tcp/4001/p2p/16Uiu2HAm7QD77kxSKf1GTM3YkrYp8vkhUwS2ySJPht9jALeaHaft",
	"/ip4/43.204.199.237/tcp/4001/p2p/16Uiu2HAm3tpaz9zgqB4i2FEwX7dwTJzv88Krpdy3kRecXZos3WdM",
	"/ip4/35.155.192.241/tcp/4001/p2p/16Uiu2HAm29iAxcKRPNRBVMYCz455uck5o7KmdPJ9GQ5BKvpxxca9",
	"/ip4/35.83.203.96/tcp/4001/p2p/16Uiu2HAmNnKCdkBKdoPo4sXSLhDgXvPmCi7NCjo8cfcP5RRb4mKL",
	"/ip4/54.69.57.58/tcp/4001/p2p/16Uiu2HAmNDZWZtyRNZMLQ88933SFcVp2gtb99aQVbADXcCFcjFn9",
	"/ip4/35.164.151.55/tcp/4001/p2p/16Uiu2HAmMgufksaU9aaenq2bNtGnG5QokCS1xdzJwUS6yRtakhbs",
}
var DefaultTestnetBootstrapAddresses = []string{
	"/ip4/18.224.174.215/tcp/45301/p2p/16Uiu2HAmFFwNdgSoLhfgJUPEfPEVodppRxaeZBVpAvrH5s3qSkWo",
	"/ip4/18.224.174.215/tcp/34237/p2p/16Uiu2HAmDigS3SDx6g9Sp6MUfdFHvDwS8Zw8E14V6bLhCAHA3jjB",
	"/ip4/18.224.174.215/tcp/43097/p2p/16Uiu2HAm7HQoEbQe1fYt4LtnG6z5TqwTrrqUv5xsnt4nukskWmAi",
	"/ip4/18.224.174.215/tcp/38955/p2p/16Uiu2HAm5WrYvkJwaRP7ZAroWCfjaUxKkNssqcSmEmKJ8vXVYp1o",
	"/ip4/54.151.185.243/tcp/36707/p2p/16Uiu2HAmDis3wAorW46YyNmXNk963VAAHwZ1phjHXj5yduyawAUy",
	"/ip4/54.151.185.243/tcp/42741/p2p/16Uiu2HAmSfqLCyqH5qQQF8zpzPMQvWiQunhWpYtSxwGw5QR2jhgU",
	"/ip4/54.151.185.243/tcp/37403/p2p/16Uiu2HAmBHwyRUETsGqjYpgPRpnMC9y39tcVYH6vKxZidCBcBeFG",
	"/ip4/54.151.185.243/tcp/37739/p2p/16Uiu2HAm2oKy37KvYmiv1nnRWZwUoLPZumNKFxPzhM1t8F3KxADu",
	"/ip4/18.158.67.141/tcp/40155/p2p/16Uiu2HAmTMEqndByECXuxk1Rg8szxMqwS3tUFFWhAUduFzwfwmfK",
	"/ip4/18.158.67.141/tcp/44569/p2p/16Uiu2HAmL4QNi68nSNbedUWp1A1cRR3z3NuJqQYmAYoj19ht6iNv",
	"/ip4/18.158.67.141/tcp/39703/p2p/16Uiu2HAkzF6JMx4EL2C4cLoCLyQH8t1sgyttQxPfQtNt5FZhvpxs",
	"/ip4/18.158.67.141/tcp/46713/p2p/16Uiu2HAm85HXJA7xmgNxxTVdFRuRCGstvrY8nW6KqfTtkuZrZg64",
	"/ip4/18.163.235.175/tcp/36335/p2p/16Uiu2HAm8wVUsVsqksBfxy6yzHpVv5gELQnpU7Q2uhDyXFwr9bfV",
	"/ip4/18.163.235.175/tcp/44029/p2p/16Uiu2HAmBvnQU5FWgEcfY1jaAK2Q9iQBy6FwQdDUtyT7mo8HU1Yu",
	"/ip4/18.163.235.175/tcp/40191/p2p/16Uiu2HAkurshicwtTrqbrL3yv9xR7hogPvreUHJP3W8n9W5XMibz",
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
Please report it to https://github.com/bittorrent/go-btfs/issues.`, err)
	}
	return ps, nil
}
func DefaultTestnetBootstrapPeers() ([]peer.AddrInfo, error) {
	ps, err := ParseBootstrapPeers(DefaultTestnetBootstrapAddresses)
	if err != nil {
		return nil, fmt.Errorf(`Failed to parse hardcoded testnet bootstrap peers: %s
This is a problem with the BTFS codebase.
Please report it to https://github.com/bittorrent/go-btfs/issues.`, err)
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
