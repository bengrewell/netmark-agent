package core

import (
	"github.com/bengrewell/netmark-agent/core/cpu"
	"github.com/bengrewell/netmark-agent/core/disk"
	"github.com/bengrewell/netmark-agent/core/docker"
	"github.com/bengrewell/netmark-agent/core/host"
	"github.com/bengrewell/netmark-agent/core/memory"
	"github.com/bengrewell/netmark-agent/core/network"
	"github.com/bengrewell/netmark-agent/core/process"
	"github.com/bengrewell/netmark-agent/core/users"
)

var (
	CoreEngine Engine
)

func init() {
	CoreEngine = Engine{Modules: []Module{
		&cpu.InfoCpu{},
		&disk.InfoDisk{},
		&docker.InfoDocker{},
		&host.InfoHost{},
		&memory.InfoMemory{},
		&network.InfoNetwork{},
		&process.InfoProcess{},
		&users.InfoUsers{},
	}}
}

type Engine struct {
	Modules []Module
}

func (e *Engine) Run() (collector map[string]interface{}) {
	collector = make(map[string]interface{})
	for _, mod := range e.Modules {
		v := mod.Run()
		for key, value := range v {
			collector[key] = value
		}
	}
	return collector
}
