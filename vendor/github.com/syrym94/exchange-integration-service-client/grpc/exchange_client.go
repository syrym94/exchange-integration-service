package grpc

import (
	"context"
	"github.com/syrym94/exchange-integration-service-client/proto"
	"google.golang.org/grpc"
)

type ExchangeClient struct {
	client proto.ExchangeServiceClient
}

func NewExchangeClient(conn *grpc.ClientConn) *ExchangeClient {
	c := new(ExchangeClient)

	c.client = proto.NewExchangeServiceClient(conn)

	return c
}

func (ec *ExchangeClient) GetTrades(exchange string) ([]*proto.Trade, error) {
	req := &proto.GetTradesRequest{Exchange: exchange}
	res, err := ec.client.GetTrades(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return res.Trades, nil
}

func (ec *ExchangeClient) GetWalletBalance(exchange, accountType string) (*proto.WalletBalanceResponse, error) {
	req := &proto.GetWalletBalanceRequest{AccountType: accountType, Exchange: exchange}
	res, err := ec.client.GetWalletBalance(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return res.Balance, nil
}

func (ec *ExchangeClient) StreamTicker(exchange, tickerSymbol string) (*proto.TickerResponse, error) {
	req := &proto.TickerRequest{Exchange: exchange, TickerSymbol: tickerSymbol}
	res, err := ec.client.StreamTickerData(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return res.Recv()
}

func (ec *ExchangeClient) GetSubDepositAddress(exchange, coin, chainType, subMemberId string) (*proto.GetSubDepositAddressResponse, error) {
	req := &proto.GetSubDepositAddressRequest{Exchange: exchange, Coin: coin, ChainType: chainType, SubMemberId: subMemberId}
	res, err := ec.client.GetSubDepositAddress(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (ec *ExchangeClient) GetAccountCoinsBalance(exchange, memberId, accountType, coin string, withBonus int) (*proto.AccountCoinsBalanceResponse, error) {
	req := &proto.AccountCoinsBalanceRequest{Exchange: exchange, Coin: coin, MemberId: memberId, AccountType: accountType, WithBonus: int64(withBonus)}
	res, err := ec.client.GetAccountCoinsBalance(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (ec *ExchangeClient) GetWithdrawalRecords(exchange, coin, withdrawId, cursor string, withdrawType, limit int32, startTime, endTime, withBonus int64) (*proto.WithdrawalRecordsResponse, error) {
	req := &proto.WithdrawalRecordsRequest{Exchange: exchange, Coin: coin, WithdrawId: withdrawId, Cursor: cursor, WithdrawType: withdrawType, Limit: limit, StartTime: startTime, EndTime: endTime}
	res, err := ec.client.GetWithdrawalRecords(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (ec *ExchangeClient) GetWithdrawableAmount(exchange, coin string) (*proto.WithdrawableAmountResponse, error) {
	req := &proto.WithdrawableAmountRequest{Exchange: exchange, Coin: coin}
	res, err := ec.client.GetWithdrawableAmount(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (ec *ExchangeClient) CreateWithdrawal(exchange, coin, chain, address, tag, accountType, amount string, timestamp int64, forceChain int32) (*proto.CreateWithdrawalResponse, error) {
	req := &proto.CreateWithdrawalRequest{Exchange: exchange, Coin: coin, Chain: chain, Address: address, Tag: tag, Amount: amount, Timestamp: timestamp, ForceChain: forceChain, AccountType: accountType}
	res, err := ec.client.CreateWithdrawal(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (ec *ExchangeClient) GetSubWithdrawalRecords(exchange, subMemberId, coin, cursor string, limit int32, startTime, endTime int64) (*proto.SubDepositRecordsResponse, error) {
	req := &proto.SubDepositRecordsRequest{Exchange: exchange, SubMemberId: subMemberId, Coin: coin, Cursor: cursor, Limit: limit, StartTime: startTime, EndTime: endTime}
	res, err := ec.client.GetSubDepositRecords(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return res, err
}
