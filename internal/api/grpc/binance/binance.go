package binance

import (
	"context"
	"github.com/syrym94/exchange-integration-service-client/proto"
	"github.com/syrym94/exchange-integration-service/internal/config"
	"github.com/syrym94/exchange-integration-service/internal/services/binance"
	"strconv"
)

type BinanceGRPC struct {
	service *binance.BinanceService
	cfg     *config.Config
}

func (s *BinanceGRPC) GetSubDepositRecords(ctx context.Context, req *proto.SubDepositRecordsRequest) (*proto.SubDepositRecordsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *BinanceGRPC) CreateWithdrawal(ctx context.Context, req *proto.CreateWithdrawalRequest) (*proto.CreateWithdrawalResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *BinanceGRPC) GetWithdrawableAmount(ctx context.Context, req *proto.WithdrawableAmountRequest) (*proto.WithdrawableAmountResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *BinanceGRPC) GetWithdrawalRecords(ctx context.Context, req *proto.WithdrawalRecordsRequest) (*proto.WithdrawalRecordsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *BinanceGRPC) StreamTickerData(req *proto.TickerRequest, stream proto.ExchangeService_StreamTickerDataServer) error {
	//TODO implement me
	panic("implement me")
}

func (s *BinanceGRPC) GetSubDepositAddress(ctx context.Context, req *proto.GetSubDepositAddressRequest) (*proto.GetSubDepositAddressResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *BinanceGRPC) GetAccountCoinsBalance(ctx context.Context, req *proto.AccountCoinsBalanceRequest) (*proto.AccountCoinsBalanceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewBinanceGRPC(service *binance.BinanceService, cfg *config.Config) *BinanceGRPC {
	return &BinanceGRPC{service: service, cfg: cfg}
}

func (s *BinanceGRPC) GetWalletBalance(ctx context.Context, req *proto.GetWalletBalanceRequest) (*proto.GetWalletBalanceResponse, error) {
	walletBalance, err := s.service.GetWalletBalance(req.AccountType, s.cfg)
	if err != nil {
		return nil, err
	}

	var grpcWalletBalanceLists []*proto.WalletBalanceList
	for _, l := range walletBalance.Result.List {
		var grpcWalletBalanceCoins []*proto.WalletBalanceCoin
		for _, c := range l.Coin {
			grpcWalletBalanceCoins = append(grpcWalletBalanceCoins, &proto.WalletBalanceCoin{
				AvailableToBorrow:   c.AvailableToBorrow,
				Bonus:               c.Bonus,
				AccruedInterest:     c.AccruedInterest,
				AvailableToWithdraw: c.AvailableToWithdraw,
				TotalOrderIM:        c.TotalOrderIM,
				Equity:              c.Equity,
				TotalPositionMM:     c.TotalPositionMM,
				UsdValue:            c.UsdValue,
				UnrealisedPnl:       c.UnrealisedPnl,
				CollateralSwitch:    c.CollateralSwitch,
				SpotHedgingQty:      c.SpotHedgingQty,
				BorrowAmount:        c.BorrowAmount,
				TotalPositionIM:     c.TotalPositionIM,
				WalletBalance:       c.WalletBalance,
				CumRealisedPnl:      c.CumRealisedPnl,
				Locked:              c.Locked,
				MarginCollateral:    c.MarginCollateral,
				Coin:                c.Coin,
			})
		}
		grpcWalletBalanceLists = append(grpcWalletBalanceLists, &proto.WalletBalanceList{
			TotalEquity:            l.TotalEquity,
			AccountIMRate:          l.AccountIMRate,
			TotalMarginBalance:     l.TotalMarginBalance,
			TotalInitialMargin:     l.TotalInitialMargin,
			AccountType:            l.AccountType,
			TotalAvailableBalance:  l.TotalAvailableBalance,
			AccountMMRate:          l.AccountMMRate,
			TotalPerpUPL:           l.TotalPerpUPL,
			TotalWalletBalance:     l.TotalWalletBalance,
			AccountLTV:             l.AccountLTV,
			TotalMaintenanceMargin: l.TotalMaintenanceMargin,
			Coin:                   grpcWalletBalanceCoins,
		})
	}

	return &proto.GetWalletBalanceResponse{
		Balance: &proto.WalletBalanceResponse{
			CommonResponse: &proto.CommonResponse{
				RetCode: int32(walletBalance.RetCode),
				RetMsg:  walletBalance.RetMsg,
				Time:    strconv.FormatInt(walletBalance.Time, 10),
			},
			List: grpcWalletBalanceLists,
		},
	}, nil
}

func (s *BinanceGRPC) GetTrades(ctx context.Context, req *proto.GetTradesRequest) (*proto.GetTradesResponse, error) {
	trades, err := s.service.GetTrades(req.Exchange, s.cfg)
	if err != nil {
		return nil, err
	}

	var grpcTrades []*proto.Trade
	for _, trade := range trades {
		grpcTrades = append(grpcTrades, &proto.Trade{
			Id:       trade.Id,
			Price:    trade.Price,
			Quantity: trade.Quantity,
			Symbol:   trade.Symbol,
			Time:     trade.Time,
		})
	}

	return &proto.GetTradesResponse{Trades: grpcTrades}, nil
}
