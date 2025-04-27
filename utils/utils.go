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
	choice := 19
	if choice >= len(ifaces) {
		fmt.Printf("%d\n", choice)
	}

	for _, iface := range ifaces {
		fmt.Printf("Nom d'interface disponible sur la machine: %s\n", iface.Name)

	}

}
