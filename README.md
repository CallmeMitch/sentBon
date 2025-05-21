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


## Techno/Packages

Frontend: React
Backend: Golang

Package golang utilisé :
fmt, log, net


