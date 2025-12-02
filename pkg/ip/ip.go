package ip

import "net/netip"

func StringToNetIpAddr(ip string) *netip.Addr {
	var ipAddr = netip.MustParseAddr(ip)

	return &ipAddr
}
