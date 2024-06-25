package http_server

import (
	"encoding/json"
	"github.com/syrym94/exchange-integration-service/internal/api/http/binance"
	"github.com/syrym94/exchange-integration-service/internal/api/http/bybit"
	"github.com/syrym94/exchange-integration-service/internal/config"
	"net/http"

	"github.com/gorilla/mux"
)

var cfg *config.Config

func RegisterRoutes(r *mux.Router, config *config.Config) {
	cfg = config
	r.HandleFunc("/trades/{exchange}", handleGetTrades).Methods("GET")
}

func handleGetTrades(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	exchange := vars["exchange"]

	switch exchange {
	case "binance":
		trades, err := binance.GetBinanceTrades(cfg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(trades)
	case "bybit":
		trades, err := bybit.GetBybitTrades(cfg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(trades)
	}

}
