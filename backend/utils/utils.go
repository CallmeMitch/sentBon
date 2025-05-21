package utils

import (
	"fmt"
	"net"
)

type InterfaceInfo struct {
	Name         string `json:"name"`
	HardwareAddr string `json:"mac"`
	MTU          int    `json:"mtu"`
}

func LocalAddresses() ([]InterfaceInfo, error) {

	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("localAddresses: %+v\n", err.Error())
	}

	var result []InterfaceInfo
	for _, iface := range ifaces {
		info := InterfaceInfo{
			Name:         iface.Name,
			HardwareAddr: iface.HardwareAddr.String(),
			MTU:          iface.MTU,
		}
		result = append(result, info)
	}

	return result, nil

}
