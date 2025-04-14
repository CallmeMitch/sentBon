# Routeur virtuel Golang


## Etapes

1. Sniffer les paquets
Sniffe une interface virtuelle, avec des raw sockets ou gopacket(A choisir et comprendre la différence entre les deux).

2. Analyser les trames
Lire les entêtes Ethernet, IP, TCP/UDP. Extraire la destination.

3. Prendre une décision de routage
Lire une table de routage (même simple, codée en Go)

Choisir une interface de sortie

Appliquer éventuellement du NAT (translation d’adresse)

4. Forwarder le paquet
Réinjecter le paquet manuellement (ou via raw socket) sur l'interface de sortie.


## Techno/Packages utiles

gopacket – pour capturer et analyser les paquets

raw socket (syscall ou via github.com/mdlayher/raw)

netlink – pour interagir avec la table de routage Linux (optionnel)

iptables/nftables – pour tester le comportement avec des règles système

ebpf (optionnel mais sexy) – pour hooker le kernel (plus complexe)