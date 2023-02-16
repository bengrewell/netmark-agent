package memory

import (
	"github.com/shirou/gopsutil/mem"
)

type InfoMemory struct {
	Memory *mem.VirtualMemoryStat `json:"memory"`
}

func (i *InfoMemory) Run() (value map[string]interface{}) {
	key := "memory"
	v, err := mem.VirtualMemory()
	if err != nil {
		return map[string]interface{}{key: err}
	}

	return map[string]interface{}{key: v}
}
