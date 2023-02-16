package network

import (
	"github.com/shirou/gopsutil/net"
)

type InfoNetwork struct {
	Network []net.InterfaceStat `json:"network"`
}

func (i *InfoNetwork) Run() (value map[string]interface{}) {
	key := "network"
	v, err := net.Interfaces()
	if err != nil {
		return map[string]interface{}{key: err}
	}

	return map[string]interface{}{key: v}
}
