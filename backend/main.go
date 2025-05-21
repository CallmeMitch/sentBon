package main

import (
	"log"
	"net/http"
	"sentbon/backend/api"
)

func main() {

	// Route HTTP qui déclenche ta fonction PrintInterfaces
	http.HandleFunc("/api/interfaces", api.PrintInterfaces)

	// Démarrage du serveur sur le port 8080
	log.Println("Serveur démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
