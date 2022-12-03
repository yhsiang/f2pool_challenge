package models

import "net"

func NewModelAddresses(addrs []net.IP) (output []*ModelAddress) {
	for _, addr := range addrs {
		output = append(output, &ModelAddress{
			IP: addr.String(),
		})
	}
	return
}
