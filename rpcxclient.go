package rpcxclient

import (
	"context"
	"github.com/carefreex-io/config"
	"github.com/smallnest/rpcx/client"
	"time"
)

type Client struct {
	XClient       client.XClient
	BaseOptions   *BaseOptions
	CustomOptions *CustomOptions
}

type BaseOptions struct {
	RegistryOption RegistryOption
	Timeout        time.Duration
}

type RegistryOption struct {
	Addr  []string
	Group string
}

type CustomOptions struct {
	BasePath   string
	ServerName string
	Breaker    func() client.Breaker
}

var (
	baseOptions          *BaseOptions
	DefaultCustomOptions = &CustomOptions{
		Breaker: func() client.Breaker {
			return client.NewConsecCircuitBreaker(3, 30*time.Second)
		},
	}
)

func initBaseOptions() {
	baseOptions = &BaseOptions{
		RegistryOption: RegistryOption{
			Addr:  config.GetStringSlice("Registry.Addr"),
			Group: config.GetString("Registry.Group"),
		},
		Timeout: config.GetDuration("Rpc.WithTimout") * time.Second,
	}
}

func NewClient() (c *Client, err error) {
	initBaseOptions()

	c = &Client{
		BaseOptions:   baseOptions,
		CustomOptions: DefaultCustomOptions,
	}
	xClient, err := c.newRpcXClient()
	c.XClient = xClient

	return c, err
}

func (c *Client) Call(ctx context.Context, method string, request interface{}, response interface{}) (err error) {
	if c.BaseOptions.Timeout > 0 {
		return c.callWithTimeout(ctx, method, request, response)
	}
	return c.XClient.Call(ctx, method, request, response)
}

func (c *Client) callWithTimeout(ctx context.Context, method string, request interface{}, response interface{}) (err error) {
	ctx, cancel := context.WithTimeout(ctx, c.BaseOptions.Timeout)
	defer cancel()

	return c.XClient.Call(ctx, method, request, response)
}
