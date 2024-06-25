package grpc

import (
	"github.com/pkg/errors"
	client "github.com/syrym94/exchange-integration-service-client"
	"google.golang.org/grpc"
)

type Client struct {
	address     string
	conn        *grpc.ClientConn
	dialOpts    []grpc.DialOption
	middlewares []grpc.UnaryClientInterceptor
}

func NewClient(opts ...Option) (*Client, error) {
	cli := &Client{
		middlewares: make([]grpc.UnaryClientInterceptor, 0),
		dialOpts:    make([]grpc.DialOption, 0),
	}

	for _, opt := range opts {
		opt(cli)
	}

	if err := cli.connect(); err != nil {
		return nil, err
	}

	return cli, nil
}

func (c *Client) connect() error {
	if len(c.middlewares) == 1 {
		c.dialOpts = append(c.dialOpts, grpc.WithUnaryInterceptor(c.middlewares[0]))
	} else if len(c.middlewares) > 1 {
		c.dialOpts = append(c.dialOpts, grpc.WithChainUnaryInterceptor(c.middlewares...))
	}

	conn, err := grpc.Dial(c.address, c.dialOpts...)
	if err != nil {
		return errors.Wrap(err, "[ERROR] Error establishing protobuf connection to jam seo")
	}

	c.conn = conn

	return nil
}

func (c *Client) ExchangeClient() client.ExchangeClient {
	return NewExchangeClient(c.conn)
}

func (c *Client) Close() error {
	return c.conn.Close()
}
