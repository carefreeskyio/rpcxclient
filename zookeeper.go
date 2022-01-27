// +build zookeeper

package rpcxclient

import (
	"github.com/smallnest/rpcx/client"
)

func (c *Client) newRpcXClient() (cli client.XClient, err error) {
	d, err := client.NewZookeeperDiscovery(c.Options.RegistryOption.BasePath, c.Options.RegistryOption.ServerName, c.Options.RegistryOption.Addr, false, nil)
	if err != nil {
		return nil, err
	}
	option := client.DefaultOption
	if c.Options.Breaker != nil {
		option.GenBreaker = c.Options.Breaker
	}
	option.Group = c.Options.RegistryOption.Group
	cli = client.NewXClient(c.Options.RegistryOption.ServerName, client.Failtry, client.RandomSelect, d, option)

	return cli, nil
}
