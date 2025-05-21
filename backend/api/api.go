package api

import (
	"encoding/json"
	"net/http"
	"sentbon/backend/utils"
)

func PrintInterfaces(w http.ResponseWriter, r *http.Request) {
	interfaces, err := utils.LocalAddresses()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(interfaces)
}
