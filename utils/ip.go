package utils

import "net/netip"

func StringToNetIpAddr(ip string) *netip.Addr {
	var ipAdd = netip.MustParseAddr(ip)

	return &ipAdd
}
