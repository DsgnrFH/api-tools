package handlers

import (
	"encoding/json"
	"net/http"
	"dsgnrfh/pkg/network"
)

func SubnetHandler(w http.ResponseWriter, r *http.Request) {
	cidr := r.URL.Query().Get("cidr")
	if cidr == "" {
		cidr = "192.168.1.0/24"
	}

	details, err := network.CalculateSubnet(cidr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid CIDR block format"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(details)
}
