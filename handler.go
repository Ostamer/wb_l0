package main

import (
	"encoding/json"
	"net/http"
)

func getOrderHandler(w http.ResponseWriter, r *http.Request) {
	orderUID := r.URL.Query().Get("id")
	if orderUID == "" {
		http.Error(w, "ID parameter is missing", http.StatusBadRequest)
		return
	}

	cacheLock.RLock()
	order, ok := cache[orderUID]
	cacheLock.RUnlock()

	if !ok {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}
