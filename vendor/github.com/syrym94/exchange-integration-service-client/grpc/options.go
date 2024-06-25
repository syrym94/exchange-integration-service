package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Option func(*Client)

func WithAddress(address string) Option {
	return func(c *Client) {
		c.address = address
	}
}

func WithInsecure() Option {
	return func(c *Client) {
		c.dialOpts = append(c.dialOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
}
