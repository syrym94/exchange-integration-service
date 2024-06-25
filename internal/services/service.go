package services

import (
	"fmt"
	"github.com/syrym94/exchange-integration-service/internal/config"
	"github.com/syrym94/exchange-integration-service/internal/models"
)

type ExchangeServicer interface {
	GetTrades(exchange string, cfg *config.Config) ([]models.Trade, error)
	GetWalletBalance(accountType string, cfg *config.Config) (*models.WalletBalance, error)
	GetSubDepositAddress(coin, chainType, subMemberId string, cfg *config.Config) (*models.DepositAddressResponse, error)
	GetAccountCoinsBalance(memberId, accountType, coin string, withBonus int, cfg *config.Config) (*models.AccountCoinBalance, error)
	GetWithdrawalRecords(coin, withdrawId, cursor string, withdrawType, limit int32, startTime, endTime int64, cfg *config.Config) (*models.WithdrawalRecords, error)
	GetWithdrawableAmount(coin string, cfg *config.Config) (*models.WithdrawableAmount, error)
	CreateWithdrawal(coin, chain, address, tag, accountType, amount string, timestamp int64, forceChain int32, cfg *config.Config) (string, error)
}

type Service struct {
	exchanges map[string]ExchangeServicer
}

func NewService(binance ExchangeServicer, bybit ExchangeServicer) *Service {
	return &Service{
		exchanges: map[string]ExchangeServicer{
			"binance": binance,
			"bybit":   bybit,
		},
	}
}

func (s *Service) GetTrades(exchange string, cfg *config.Config) ([]models.Trade, error) {
	service, ok := s.exchanges[exchange]
	if !ok {
		return nil, fmt.Errorf("unsupported exchange: %s", exchange)
	}
	return service.GetTrades(exchange, cfg)
}

func (s *Service) GetWalletBalance(exchange, accountType string, cfg *config.Config) (*models.WalletBalance, error) {
	service, ok := s.exchanges[exchange]
	if !ok {
		return nil, fmt.Errorf("unsupported exchange: %s", exchange)
	}
	return service.GetWalletBalance(accountType, cfg)
}

func (s *Service) GetSubDepositAddress(exchange, coin, chainType, subMemberId string, cfg *config.Config) (*models.DepositAddressResponse, error) {
	service, ok := s.exchanges[exchange]
	if !ok {
		return nil, fmt.Errorf("unsupported exchange: %s", exchange)
	}
	return service.GetSubDepositAddress(coin, chainType, subMemberId, cfg)
}

func (s *Service) GetAccountCoinsBalance(exchange, memberId, accountType, coin string, withBonus int, cfg *config.Config) (*models.AccountCoinBalance, error) {
	service, ok := s.exchanges[exchange]
	if !ok {
		return nil, fmt.Errorf("unsupported exchange: %s", exchange)
	}
	return service.GetAccountCoinsBalance(memberId, coin, accountType, withBonus, cfg)
}

func (s *Service) GetWithdrawalRecords(exchange, coin, withdrawId, cursor string, withdrawType, limit int32, startTime, endTime int64, cfg *config.Config) (*models.WithdrawalRecords, error) {
	service, ok := s.exchanges[exchange]
	if !ok {
		return nil, fmt.Errorf("unsupported exchange: %s", exchange)
	}
	return service.GetWithdrawalRecords(coin, withdrawId, cursor, withdrawType, limit, startTime, endTime, cfg)
}

func (s *Service) GetWithdrawableAmount(exchange, coin, withdrawId, cursor string, withdrawType, limit int32, startTime, endTime int64, cfg *config.Config) (*models.WithdrawableAmount, error) {
	service, ok := s.exchanges[exchange]
	if !ok {
		return nil, fmt.Errorf("unsupported exchange: %s", exchange)
	}
	return service.GetWithdrawableAmount(coin, cfg)
}

func (s *Service) CreateWithdrawal(exchange, coin, chain, address, tag, accountType, amount string, timestamp int64, forceChain int32, cfg *config.Config) (string, error) {
	service, ok := s.exchanges[exchange]
	if !ok {
		return "", fmt.Errorf("unsupported exchange: %s", exchange)
	}
	return service.CreateWithdrawal(coin, chain, address, tag, accountType, amount, timestamp, forceChain, cfg)
}
