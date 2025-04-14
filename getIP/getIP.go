package getIP

import (
	"fmt"
	"log"
	"net"
)

func ReadNIC() {

	iface := "wlp0s20f3"

	fmt.Printf("%v:\n", iface)

	netInterface, err := net.InterfaceByName(iface)
	if err != nil {
		log.Fatalf("Erreur lors de la récupération de l'interface %s: %v", iface, err)
	}

	addrs, err := netInterface.Addrs()
	if err != nil {
		log.Fatalf("Erreur lors de la récupération des adresses : %v", err)
	}

	if len(addrs) == 0 {
		fmt.Println("Aucune adresse unicast trouvée pour cette interface.")
	}

	for _, addr := range addrs {

		ipNet, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}

		ipv4 := ipNet.IP.To4()
		if ipv4 != nil {
			fmt.Println(ipv4.String())
		}

		conn, err := net.ListenPacket("ip4:tcp", ipv4.String())
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		buffer := make([]byte, 65536)

		for {
			n, _, err := conn.ReadFrom(buffer)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(buffer[:n])
		}
	}

}
