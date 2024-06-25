package bybit

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/syrym94/exchange-integration-service-client/proto"
	"github.com/syrym94/exchange-integration-service/internal/config"
	"github.com/syrym94/exchange-integration-service/internal/services/bybit"
	"log"
	"strconv"
)

type BybitGRPC struct {
	service *bybit.BybitService
	cfg     *config.Config
}

func NewBybitGRPC(service *bybit.BybitService, cfg *config.Config) *BybitGRPC {
	return &BybitGRPC{service: service, cfg: cfg}
}

func (s *BybitGRPC) GetWalletBalance(ctx context.Context, req *proto.GetWalletBalanceRequest) (*proto.GetWalletBalanceResponse, error) {
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

func (s *BybitGRPC) GetTrades(ctx context.Context, req *proto.GetTradesRequest) (*proto.GetTradesResponse, error) {
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

func (s *BybitGRPC) GetSubDepositAddress(ctx context.Context, req *proto.GetSubDepositAddressRequest) (*proto.GetSubDepositAddressResponse, error) {
	res, err := s.service.GetSubDepositAddress(req.Coin, req.ChainType, req.SubMemberId, s.cfg)
	if err != nil {
		return nil, err
	}

	return &proto.GetSubDepositAddressResponse{
		Address: res.Result.Chains.AddressDeposit,
		Tag:     res.Result.Chains.TagDeposit,
	}, nil
}

func (s *BybitGRPC) GetAccountCoinsBalance(ctx context.Context, req *proto.AccountCoinsBalanceRequest) (*proto.AccountCoinsBalanceResponse, error) {
	res, err := s.service.GetAccountCoinsBalance(req.MemberId, req.Coin, req.AccountType, int(req.WithBonus), s.cfg)
	if err != nil {
		return nil, err
	}

	var coinBalances []*proto.CoinBalance
	for _, balance := range res.Result.Balance {
		coinBalances = append(coinBalances, &proto.CoinBalance{
			Coin:            balance.Coin,
			WalletBalance:   balance.WalletBalance,
			TransferBalance: balance.TransferBalance,
			Bonus:           balance.Bonus,
		})
	}

	return &proto.AccountCoinsBalanceResponse{
		CoinBalances: coinBalances,
	}, nil
}

func (s *BybitGRPC) GetWithdrawalRecords(ctx context.Context, req *proto.WithdrawalRecordsRequest) (*proto.WithdrawalRecordsResponse, error) {
	res, err := s.service.GetWithdrawalRecords(req.Coin, req.WithdrawId, req.Cursor, req.WithdrawType, req.Limit, req.StartTime, req.EndTime, s.cfg)
	if err != nil {
		return nil, err
	}

	var grpcWithdrawalRecords []*proto.WithdrawalRecord
	for _, record := range res.Result.Rows {
		grpcWithdrawalRecords = append(grpcWithdrawalRecords, &proto.WithdrawalRecord{
			Coin:         record.Coin,
			Chain:        record.Chain,
			Amount:       record.Amount,
			TxID:         record.TxID,
			Status:       record.Status,
			ToAddress:    record.ToAddress,
			Tag:          record.Tag,
			WithdrawFee:  record.WithdrawFee,
			CreateTime:   record.CreateTime,
			UpdateTime:   record.UpdateTime,
			WithdrawId:   record.WithdrawID,
			WithdrawType: int32(record.WithdrawType),
		})
	}

	return &proto.WithdrawalRecordsResponse{
		Rows: grpcWithdrawalRecords,
	}, nil
}

func (s *BybitGRPC) GetWithdrawableAmount(ctx context.Context, req *proto.WithdrawableAmountRequest) (*proto.WithdrawableAmountResponse, error) {
	res, err := s.service.GetWithdrawableAmount(req.Coin, s.cfg)
	if err != nil {
		return nil, err
	}

	grpcWithdrawableAmounts := make(map[string]*proto.WithdrawableAmount)
	grpcWithdrawableAmounts["FUND"] = &proto.WithdrawableAmount{
		Coin:               res.Result.WithdrawableAmount.FUND.Coin,
		WithdrawableAmount: res.Result.WithdrawableAmount.FUND.WithdrawableAmount,
		AvailableBalance:   res.Result.WithdrawableAmount.FUND.AvailableBalance,
	}

	grpcWithdrawableAmounts["SPOT"] = &proto.WithdrawableAmount{
		Coin:               res.Result.WithdrawableAmount.SPOT.Coin,
		WithdrawableAmount: res.Result.WithdrawableAmount.SPOT.WithdrawableAmount,
		AvailableBalance:   res.Result.WithdrawableAmount.SPOT.AvailableBalance,
	}

	return &proto.WithdrawableAmountResponse{
		LimitAmountUsd:      res.Result.LimitAmountUsd,
		WithdrawableAmounts: grpcWithdrawableAmounts,
	}, nil
}

func (s *BybitGRPC) CreateWithdrawal(ctx context.Context, req *proto.CreateWithdrawalRequest) (*proto.CreateWithdrawalResponse, error) {
	res, err := s.service.CreateWithdrawal(req.Exchange, req.Coin, req.Chain, req.Address, req.Tag, req.AccountType, req.Amount, req.Timestamp, req.ForceChain, s.cfg)
	if err != nil {
		return nil, err
	}

	return &proto.CreateWithdrawalResponse{
		Id: res,
	}, nil
}

func (s *BybitGRPC) StreamTickerData(req *proto.TickerRequest, stream proto.ExchangeService_StreamTickerDataServer) error {
	conn, _, err := websocket.DefaultDialer.Dial(s.cfg.BybitWSEndpoint, nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	subscribeMessage := map[string]interface{}{
		"op":   "subscribe",
		"args": []string{fmt.Sprintf("tickers.%s", req.TickerSymbol)},
	}
	if err := conn.WriteJSON(subscribeMessage); err != nil {
		return err
	}

	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Printf("Error reading message: %v", err)
				return
			}
			if err := stream.Send(&proto.TickerResponse{Data: string(message)}); err != nil {
				log.Printf("Error sending message: %v", err)
				return
			}
		}
	}()

	<-done
	return nil
}
