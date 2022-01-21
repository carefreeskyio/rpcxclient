//+build etcd

package rpcxclient

import (
	etcdClient "github.com/rpcxio/rpcx-etcd/client"
	"github.com/smallnest/rpcx/client"
)

func (c *Client) newRpcXClient() (cli client.XClient, err error) {
	d, err := etcdClient.NewEtcdDiscovery(c.Options.BasePath, c.Options.ServerName, c.Options.Addr, false, nil)
	if err != nil {
		return nil, err
	}
	option := client.DefaultOption
	if c.Options.EnableBreaker {
		option.GenBreaker = func() client.Breaker {
			return client.NewConsecCircuitBreaker(c.Options.FailureThreshold, c.Options.Window)
		}
	}
	option.Group = c.Options.Group
	cli = client.NewXClient(c.Options.ServerName, client.Failtry, client.RandomSelect, d, option)

	return cli, nil
}
