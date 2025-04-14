package utils

import (
	"fmt"
	"net"
)

func LocalAddresses() {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("localAddresses: %+v\n", err.Error())
		return
	}

	ipv4, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
	}

	for _, iface := range ifaces {
		net := iface.Name
		if net == "wlp0s20f3" {
			fmt.Println(net)
		} else {
			continue
		}

		for _, IP := range ipv4 {
			//for _, parsedIP := range net.ParseCIDR(IP) {
			//	fmt.Println(parsedIP)
			//}

			//ipParsed := net.ParseCIDR(IP)
			fmt.Println(IP)
		}
	}
}
