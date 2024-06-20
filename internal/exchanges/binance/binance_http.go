package binance

import (
	"encoding/json"
	"exchange-integration-service/internal/config"
	"exchange-integration-service/internal/models"
	"exchange-integration-service/internal/utils"
	"fmt"
)

func GetBinanceTrades(cfg *config.Config) ([]models.Trade, error) {
	url := fmt.Sprintf("%s/trades?symbol=BTCUSDT", cfg.BinanceAPIEndpoint)
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
