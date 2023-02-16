package cpu

import (
	"github.com/shirou/gopsutil/cpu"
)

type InfoCpu struct {
	CPU []cpu.InfoStat `json:"cpu"`
}

func (i *InfoCpu) Run() (value map[string]interface{}) {
	key := "cpu"
	v, err := cpu.Info()
	if err != nil {
		return map[string]interface{}{key: err}
	}

	return map[string]interface{}{key: v}
}
