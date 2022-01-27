package rpcxclient

import (
	"context"
	"github.com/smallnest/rpcx/client"
	"time"
)

type Client struct {
	XClient client.XClient
	Options *Options
}

type Options struct {
	RegistryOption RegistryOption
	Timeout        time.Duration
	Breaker        func() client.Breaker
}

type RegistryOption struct {
	BasePath   string
	ServerName string
	Addr       []string
	Group      string
}

var (
	DefaultOptions = &Options{
		Breaker: func() client.Breaker {
			return client.NewConsecCircuitBreaker(3, 30*time.Second)
		},
	}
)

func NewClient(options *Options) (c *Client, err error) {
	c = &Client{
		Options: options,
	}
	xClient, err := c.newRpcXClient()
	c.XClient = xClient

	return c, err
}

func (c *Client) Call(ctx context.Context, method string, request interface{}, response interface{}) (err error) {
	if c.Options.Timeout > 0 {
		return c.callWithTimeout(ctx, method, request, response)
	}
	return c.XClient.Call(ctx, method, request, response)
}

func (c *Client) callWithTimeout(ctx context.Context, method string, request interface{}, response interface{}) (err error) {
	ctx, cancel := context.WithTimeout(ctx, c.Options.Timeout)
	defer cancel()

	return c.XClient.Call(ctx, method, request, response)
}
