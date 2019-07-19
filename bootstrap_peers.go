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
	"/ip4/3.120.224.94/tcp/4001/btfs/QmWTTmvchTodUaVvuKZMo67xk7ZgkxJf4nBo7SZry3vGU5",
	"/ip4/18.194.71.27/tcp/4001/btfs/QmYHkY5CrWcvgaDo4PfvzTQgaZtfaqRGDjwW1MrHUj8cLK",
	"/ip4/54.93.47.134/tcp/4001/btfs/QmeHaHe7WvjeY37z5MYC3qYQcQcuvDwUhwTXtP3KhKLXXK",
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
		return nil, fmt.Errorf(`failed to parse hardcoded bootstrap peers: %s
This is a problem with the btfs codebase. Please report it to the dev team.`, err)
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
