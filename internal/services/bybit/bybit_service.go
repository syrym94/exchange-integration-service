package bybit

import (
	"encoding/json"
	"fmt"
	"github.com/syrym94/exchange-integration-service-client/grpc"
	"github.com/syrym94/exchange-integration-service/internal/config"
	"github.com/syrym94/exchange-integration-service/internal/models"
	"github.com/syrym94/exchange-integration-service/internal/utils"
	"log"
	"net/http"
	"strconv"
)

type BybitService struct {
	exchangeClient grpc.ExchangeClient
}

func NewBybitService() *BybitService {
	return &BybitService{}
}

func (s *BybitService) GetTrades(exchange string, config *config.Config) ([]models.Trade, error) {
	//TODO implement me
	panic("implement me")
}

func (s *BybitService) GetWalletBalance(accountType string, cfg *config.Config) (*models.WalletBalance, error) {
	url := fmt.Sprintf("%s/account/wallet-balance", cfg.BybitAPIEndpoint)
	params := map[string]interface{}{
		"accountType": accountType,
	}
	body, err := utils.MakeBybitAuthenticatedRequest(http.MethodGet, url, cfg.BybitAPIKey, cfg.BybitSecret, params)
	if err != nil {
		return nil, err
	}
	log.Println(string(body))

	var balanceResponse models.WalletBalance
	if err := json.Unmarshal(body, &balanceResponse); err != nil {
		return nil, fmt.Errorf("could not unmarshal response: %v", err)
	}

	return &balanceResponse, nil
}

func (s *BybitService) GetSubDepositAddress(coin, chainType, subMemberId string, cfg *config.Config) (*models.DepositAddressResponse, error) {
	url := fmt.Sprintf("%s/v5/asset/deposit/query-sub-member-address", cfg.BybitAPIEndpoint)
	params := map[string]interface{}{
		"coin":        coin,
		"chainType":   chainType,
		"subMemberId": subMemberId,
	}
	body, err := utils.MakeBybitAuthenticatedRequest(http.MethodGet, url, cfg.BybitAPIKey, cfg.BybitSecret, params)
	if err != nil {
		return nil, err
	}
	log.Println(string(body))

	var response models.DepositAddressResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("could not unmarshal response: %v", err)
	}

	return &response, nil
}

func (s *BybitService) GetAccountCoinsBalance(memberId, accountType, coin string, withBonus int, cfg *config.Config) (*models.AccountCoinBalance, error) {
	url := fmt.Sprintf("%s/v5/asset/deposit/query-sub-member-address", cfg.BybitAPIEndpoint)
	params := map[string]interface{}{
		"coin":        coin,
		"memberId":    memberId,
		"accountType": accountType,
		"withBonus":   strconv.Itoa(withBonus),
	}
	body, err := utils.MakeBybitAuthenticatedRequest(http.MethodGet, url, cfg.BybitAPIKey, cfg.BybitSecret, params)
	if err != nil {
		return nil, err
	}
	log.Println(string(body))

	var response models.AccountCoinBalance
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("could not unmarshal response: %v", err)
	}

	return &response, nil
}

func (s *BybitService) GetWithdrawalRecords(coin, withdrawId, cursor string, withdrawType, limit int32, startTime, endTime int64, cfg *config.Config) (*models.WithdrawalRecords, error) {
	url := fmt.Sprintf("%s/v5/asset/withdraw/query-record", cfg.BybitAPIEndpoint)
	params := map[string]interface{}{
		"coin":         coin,
		"withdrawId":   withdrawId,
		"cursor":       cursor,
		"withdrawType": strconv.Itoa(int(withdrawType)),
		"limit":        strconv.Itoa(int(limit)),
		"startTime":    strconv.Itoa(int(startTime)),
		"endTime":      strconv.Itoa(int(endTime)),
	}
	body, err := utils.MakeBybitAuthenticatedRequest(http.MethodGet, url, cfg.BybitAPIKey, cfg.BybitSecret, params)
	if err != nil {
		return nil, err
	}
	log.Println(string(body))

	var response models.WithdrawalRecords
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("could not unmarshal response: %v", err)
	}

	return &response, nil
}

func (s *BybitService) GetWithdrawableAmount(coin string, cfg *config.Config) (*models.WithdrawableAmount, error) {
	url := fmt.Sprintf("%s/v5/asset/withdraw/withdrawable-amount", cfg.BybitAPIEndpoint)
	params := map[string]interface{}{
		"coin": coin,
	}
	body, err := utils.MakeBybitAuthenticatedRequest(http.MethodGet, url, cfg.BybitAPIKey, cfg.BybitSecret, params)
	if err != nil {
		return nil, err
	}
	log.Println(string(body))

	var response models.WithdrawableAmount
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("could not unmarshal response: %v", err)
	}

	return &response, nil
}

func (s *BybitService) CreateWithdrawal(exchange, coin, chain, address, tag, accountType, amount string, timestamp int64, forceChain int32, cfg *config.Config) (string, error) {
	url := fmt.Sprintf("%s/v5/asset/withdraw/withdrawable-amount", cfg.BybitAPIEndpoint)
	params := map[string]interface{}{
		"coin":        coin,
		"chain":       chain,
		"address":     address,
		"tag":         tag,
		"amount":      amount,
		"timestamp":   timestamp,
		"forceChain":  forceChain,
		"accountType": accountType,
	}
	body, err := utils.MakeBybitAuthenticatedRequest(http.MethodPost, url, cfg.BybitAPIKey, cfg.BybitSecret, params)
	if err != nil {
		return "", err
	}
	log.Println(string(body))

	var response string
	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("could not unmarshal response: %v", err)
	}

	return response, nil
}
