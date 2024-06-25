package binance

import (
	"github.com/syrym94/exchange-integration-service-client/grpc"
	"github.com/syrym94/exchange-integration-service/internal/config"
	"github.com/syrym94/exchange-integration-service/internal/models"
)

type BinanceService struct {
	exchangeClient grpc.ExchangeClient
}

func (s *BinanceService) GetAccountCoinsBalance(memberId, accountType, coin string, withBonus int, cfg *config.Config) (*models.AccountCoinBalance, error) {
	//TODO implement me
	panic("implement me")
}

func (s *BinanceService) GetSubDepositAddress(coin, chainType, subMemberId string, cfg *config.Config) (*models.DepositAddressResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewBinanceService() *BinanceService {
	return &BinanceService{}
}

func (s *BinanceService) GetTrades(exchange string, config *config.Config) ([]models.Trade, error) {
	//TODO implement me
	panic("implement me")
}

func (s *BinanceService) GetWalletBalance(accountType string, cfg *config.Config) (*models.WalletBalance, error) {
	//TODO implement me
	panic("implement me")
}
