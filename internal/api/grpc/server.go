package grpc

import (
	"context"
	"github.com/syrym94/exchange-integration-service-client/proto"
	"github.com/syrym94/exchange-integration-service/internal/api/grpc/binance"
	"github.com/syrym94/exchange-integration-service/internal/api/grpc/bybit"
	"github.com/syrym94/exchange-integration-service/internal/config"
	binanceService "github.com/syrym94/exchange-integration-service/internal/services/binance"
	bybitService "github.com/syrym94/exchange-integration-service/internal/services/bybit"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedExchangeServiceServer
	cfg         *config.Config
	binanceGRPC *binance.BinanceGRPC
	bybitGRPC   *bybit.BybitGRPC
	exchanges   map[string]ExchangeGrpcServer
}

type ExchangeGrpcServer interface {
	GetTrades(ctx context.Context, req *proto.GetTradesRequest) (*proto.GetTradesResponse, error)
	GetWalletBalance(ctx context.Context, req *proto.GetWalletBalanceRequest) (*proto.GetWalletBalanceResponse, error)
	StreamTickerData(req *proto.TickerRequest, stream proto.ExchangeService_StreamTickerDataServer) error
	GetSubDepositAddress(ctx context.Context, req *proto.GetSubDepositAddressRequest) (*proto.GetSubDepositAddressResponse, error)
	GetAccountCoinsBalance(ctx context.Context, req *proto.AccountCoinsBalanceRequest) (*proto.AccountCoinsBalanceResponse, error)
	GetWithdrawalRecords(ctx context.Context, req *proto.WithdrawalRecordsRequest) (*proto.WithdrawalRecordsResponse, error)
	GetWithdrawableAmount(ctx context.Context, req *proto.WithdrawableAmountRequest) (*proto.WithdrawableAmountResponse, error)
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		cfg: cfg,
		exchanges: map[string]ExchangeGrpcServer{
			"binance": binance.NewBinanceGRPC(binanceService.NewBinanceService(),
				cfg),
			"bybit": bybit.NewBybitGRPC(
				bybitService.NewBybitService(),
				cfg),
		},
	}
}

func (s *Server) GetTrades(ctx context.Context, req *proto.GetTradesRequest) (*proto.GetTradesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) GetWalletBalance(ctx context.Context, req *proto.GetWalletBalanceRequest) (*proto.GetWalletBalanceResponse, error) {
	return s.exchanges[req.Exchange].GetWalletBalance(ctx, req)
}

func (s *Server) GetSubDepositAddress(ctx context.Context, req *proto.GetSubDepositAddressRequest) (*proto.GetSubDepositAddressResponse, error) {
	return s.exchanges[req.Exchange].GetSubDepositAddress(ctx, req)
}

func (s *Server) GetAccountCoinsBalance(ctx context.Context, req *proto.AccountCoinsBalanceRequest) (*proto.AccountCoinsBalanceResponse, error) {
	return s.exchanges[req.Exchange].GetAccountCoinsBalance(ctx, req)
}

func (s *Server) StreamTickerData(req *proto.TickerRequest, stream proto.ExchangeService_StreamTickerDataServer) error {
	return s.exchanges[req.Exchange].StreamTickerData(req, stream)
}

func (s *Server) GetWithdrawalRecords(ctx context.Context, req *proto.WithdrawalRecordsRequest) (*proto.WithdrawalRecordsResponse, error) {
	return s.exchanges[req.Exchange].GetWithdrawalRecords(ctx, req)
}

func (s *Server) GetWithdrawableAmount(ctx context.Context, req *proto.WithdrawableAmountRequest) (*proto.WithdrawableAmountResponse, error) {
	return s.exchanges[req.Exchange].GetWithdrawableAmount(ctx, req)
}

func StartGRPCServer(port string, cfg *config.Config) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterExchangeServiceServer(s, NewServer(cfg))

	log.Printf("gRPC server listening on %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
