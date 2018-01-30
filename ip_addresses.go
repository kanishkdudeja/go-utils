package utils

import (
	"net"
)

// Private IPV4 Address Blocks (CIDR) as per https://en.wikipedia.org/wiki/Reserved_IP_addresses
var privateIPV4Blocks = []string{
	"10.0.0.0/8",
	"100.64.0.0/10",
	"172.16.0.0/12",
	"192.0.0.0/24",
	"192.168.0.0/16",
	"198.18.0.0/15",
}

// Private IPV6 Address Blocks (CIDR) as per https://en.wikipedia.org/wiki/Reserved_IP_addresses
var privateIPV6Blocks = []string{
	"fc00::/7",
}

// inRange checks to see if a given ip address is within a IP.net
func inRange(net *net.IPNet, ipAddress net.IP) bool {
	return net.Contains(ipAddress)
}

// isPrivateSubnet function check to see if this ip is in a private subnet
func isPrivateSubnet(ipAddress net.IP) bool {
	// my use case is only concerned with ipv4 atm

	privateBlocks := []string{}

	if isIPV4(ipAddress) {
		privateBlocks = privateIPV4Blocks
	} else if isIPV6(ipAddress) {
		privateBlocks = privateIPV6Blocks
	} else {
		return false
	}

	for _, block := range privateBlocks {
		_, net, err := net.ParseCIDR(block)

		if err != nil {
			return false
		}

		if inRange(net, ipAddress) {
			return true
		}
	}

	return false
}

func isIPV4(ipAddress net.IP) bool {
	if ipAddress.To4() != nil {
		return true
	}

	return false
}

func isIPV6(ipAddress net.IP) bool {
	if ipAddress.To16() != nil {
		return true
	}

	return false
}
