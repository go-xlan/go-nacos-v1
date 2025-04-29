package utils

import (
	"net"

	"github.com/yyle88/erero"
)

func GetIpv4() (string, error) {
	nets := map[string]bool{
		"en0":    true,
		"eth0":   true,
		"ens224": true,
		"ens5":   true,
	}
	var ips []string

	interfaces, err := net.Interfaces()
	if err != nil {
		return "", erero.Wro(err)
	}

	for _, item := range interfaces {
		if item.Flags&net.FlagUp != net.FlagUp {
			continue
		}
		addresses, err := item.Addrs()
		if err != nil {
			continue // Skip this interface
		}

		// Check if the current interface is the specified one
		if nets[item.Name] {
			for _, address := range addresses {
				switch ip := address.(type) {
				case *net.IPNet:
					if !ip.IP.IsLoopback() && ip.IP.To4() != nil {
						ips = append(ips, ip.IP.String())
					}
				default:
					continue // Ignore non-IP addresses
				}
			}
		}
	}

	if len(ips) == 0 {
		return "", erero.New("没有从本地网卡找到ipv4")
	}
	return ips[0], nil
}
