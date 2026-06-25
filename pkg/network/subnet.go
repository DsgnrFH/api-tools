package network

import (
	"net"
)

type SubnetDetails struct {
	CIDR          string `json:"cidr"`
	Netmask       string `json:"netmask"`
	Broadcast     string `json:"broadcast"`
	FirstAddress  string `json:"first_usable"`
	LastAddress   string `json:"last_usable"`
	TotalHosts    uint64 `json:"total_hosts"`
}

func CalculateSubnet(cidrStr string) (*SubnetDetails, error) {
	_, ipNet, err := net.ParseCIDR(cidrStr)
	if err != nil {
		return nil, err
	}

	mask := ipNet.Mask
	netmask := net.IP(mask).String()

	// Calculate broadcast address
	broadcast := make(net.IP, len(ipNet.IP))
	for i := range ipNet.IP {
		broadcast[i] = ipNet.IP[i] | ^mask[i]
	}

	// Calculate first and last usable hosts
	firstIP := make(net.IP, len(ipNet.IP))
	copy(firstIP, ipNet.IP)
	firstIP[len(firstIP)-1]++

	lastIP := make(net.IP, len(broadcast))
	copy(lastIP, broadcast)
	lastIP[len(lastIP)-1]--

	// Calculate total hosts (2**(bits) - 2 for IPv4 networks)
	ones, bits := mask.Size()
	var totalHosts uint64
	if bits-ones < 64 {
		totalHosts = 1 << (bits - ones)
		if totalHosts > 2 {
			totalHosts -= 2 // Subtract network and broadcast
		}
	}

	return &SubnetDetails{
		CIDR:         cidrStr,
		Netmask:      netmask,
		Broadcast:    broadcast.String(),
		FirstAddress: firstIP.String(),
		LastAddress:  lastIP.String(),
		TotalHosts:   totalHosts,
	}, nil
}
