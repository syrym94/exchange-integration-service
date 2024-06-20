package http_server

import (
	"encoding/json"
	"exchange-integration-service/internal/config"
	"exchange-integration-service/internal/services"
	"net/http"

	"github.com/gorilla/mux"
)

var cfg *config.Config

func RegisterRoutes(r *mux.Router, config *config.Config) {
	cfg = config
	r.HandleFunc("/trades/{exchange}", handleGetTrades).Methods("GET")
	r.HandleFunc("/ws", handleWebSocket) // WebSocket route
}

func handleGetTrades(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	exchange := vars["exchange"]

	trades, err := services.GetTrades(exchange, cfg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(trades)
}
