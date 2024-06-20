package bybit

import (
	"encoding/json"
	"exchange-integration-service/internal/config"
	"exchange-integration-service/internal/models"
	"exchange-integration-service/internal/utils"
	"fmt"
)

func GetBybitTrades(cfg *config.Config) ([]models.Trade, error) {
	url := fmt.Sprintf("%s/trading-records?symbol=BTCUSD", cfg.BybitAPIEndpoint)
	body, err := utils.MakeHTTPRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var trades []models.Trade
	if err := json.Unmarshal(body, &trades); err != nil {
		return nil, fmt.Errorf("could not unmarshal response: %v", err)
	}

	return trades, nil
}
