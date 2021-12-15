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
}
var DefaultTestnetBootstrapAddresses = []string{
	"/ip4/18.158.67.141/tcp/4000/p2p/16Uiu2HAmNvMMxEEfP1mg5d6TUQQLWhHd4s4moAB8xPgCgZBTwwZF",
	"/ip4/18.158.67.141/tcp/4001/p2p/16Uiu2HAmUYNteRBWb12KGZ2UwogfuVdCFJfuH3fYP3bcndTUY4ZF",
	"/ip4/18.158.67.141/tcp/4002/p2p/16Uiu2HAmJCoTMZkmW1zkCVCnohTiCk1PyepxeGMgnvQYpnQMxqVb",
	"/ip4/18.158.67.141/tcp/4003/p2p/16Uiu2HAkxcrp6Pf5WvzQEEbEMVMZfTrtMpnTJCJcUaN6PGDKvegX",
	"/ip4/18.163.235.175/tcp/4000/p2p/16Uiu2HAkw8Pkrf4fARvKZUcPBH7LhhVG4PoZBqKni3kbWkYsPYog",
	"/ip4/18.163.235.175/tcp/4001/p2p/16Uiu2HAmJ9xHAr4ymFBt16tvst1psSc9jMRsbxvUDC2xw1XtnAEg",
	"/ip4/18.163.235.175/tcp/4002/p2p/16Uiu2HAmLVmEkbxbq3Pmem5qun8HB1t9npkjycmqCf3Sk4wy8EPx",
	"/ip4/18.163.235.175/tcp/4003/p2p/16Uiu2HAmLxRpzJKTzGkW8ajtbg89p8TTgagdHrvAWbum7rftTrQS",
	"/ip4/18.224.174.215/tcp/4000/p2p/16Uiu2HAmDonyuAvQtGJ7kvM5xB3aJz25niDAzYoYQUjSuaZJopoZ",
	"/ip4/18.224.174.215/tcp/4001/p2p/16Uiu2HAm3CyQ4o9adf642iZDUfQW4qVsm9FeBummYDzY9ExTrf3v",
	"/ip4/18.224.174.215/tcp/4002/p2p/16Uiu2HAkyu5b6D7TwLNSQfvK2xyZWevTaLLaiKG2tKsQfd272LGs",
	"/ip4/18.224.174.215/tcp/4003/p2p/16Uiu2HAm6biQRpXPKi92X4i9bCjnbTtPAmQ97czMjWhutvdE9m46",
	"/ip4/54.151.185.243/tcp/4000/p2p/16Uiu2HAmU9ysnuasmdyq1rRePYTwHntmyhZdfC9wm4qCPQMAh9Qq",
	"/ip4/54.151.185.243/tcp/4001/p2p/16Uiu2HAmBr4p1met83wxVCyVwpQexBVzSZrawfnLeWH9Ce8wwwtg",
	"/ip4/54.151.185.243/tcp/4002/p2p/16Uiu2HAmBqxLRBmjF16sQuyQxdWf9dgXVzY3Qp5YXjKT5zgkmsVN",
	"/ip4/54.151.185.243/tcp/4003/p2p/16Uiu2HAmSg1CHTDNEazJnDGfhA5pGbw1RZK1g2HZJmi3JVPkHGW8",
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
