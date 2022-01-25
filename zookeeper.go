// +build zookeeper

package rpcxclient

import (
	"github.com/smallnest/rpcx/client"
)

func (c *Client) newRpcXClient() (cli client.XClient, err error) {
	d, err := client.NewZookeeperDiscovery(c.CustomOptions.BasePath, c.CustomOptions.ServerName, c.BaseOptions.RegistryOption.Addr, false, nil)
	if err != nil {
		return nil, err
	}
	option := client.DefaultOption
	if c.CustomOptions.Breaker != nil {
		option.GenBreaker = c.CustomOptions.Breaker
	}
	option.Group = c.BaseOptions.RegistryOption.Group
	cli = client.NewXClient(c.CustomOptions.ServerName, client.Failtry, client.RandomSelect, d, option)

	return cli, nil
}
