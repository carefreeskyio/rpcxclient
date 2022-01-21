package rpcxclient

import (
	"context"
	"github.com/smallnest/rpcx/client"
	"time"
)

type Client struct {
	XClient client.XClient
	Options Options
}

type Options struct {
	BasePath         string
	ServerName       string
	Addr             []string
	Group            string
	Timeout          time.Duration
	FailureThreshold uint64
	Window           time.Duration
	EnableBreaker    bool
}

var DefaultOptions = Options{
	Timeout:          time.Second,
	FailureThreshold: 5,
	Window:           30 * time.Second,
	EnableBreaker:    true,
}

func NewClient(options Options) (c *Client, err error) {
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
