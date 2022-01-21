package rpcxclient

import (
	etcdClient "github.com/rpcxio/rpcx-etcd/client"
	"github.com/smallnest/rpcx/client"
)

func (c *Client) newRpcXClient() (cli client.XClient, err error) {
	d, err := etcdClient.NewEtcdV3Discovery(c.Options.BasePath, c.Options.ServerName, c.Options.Addr, false, nil)
	if err != nil {
		return nil, err
	}
	option := client.DefaultOption
	if c.Options.Breaker != nil {
		option.GenBreaker = c.Options.Breaker
	}
	option.Group = c.Options.Group
	cli = client.NewXClient(c.Options.ServerName, client.Failtry, client.RandomSelect, d, option)

	return cli, nil
}
