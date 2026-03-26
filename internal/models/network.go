package models

import (
	"net"
)

// ExpandCIDR takes a CIDR string and returns a slice of all IP addresses in the range.
func ExpandCIDR(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}

	// Remove network and broadcast addresses for IPv4 if it's not a /32 or /31
	if len(ips) > 2 && ip.To4() != nil {
		return ips[1 : len(ips)-1], nil
	}

	return ips, nil
}

// inc increments an IP address.
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
