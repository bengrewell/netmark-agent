package disk

import (
	"github.com/shirou/gopsutil/disk"
)

type InfoDisk struct {
	Disk *disk.UsageStat `json:"disk"`
}

func (i *InfoDisk) Run() (value map[string]interface{}) {
	key := "disk"
	v, err := disk.Usage("/")
	if err != nil {
		return map[string]interface{}{key: err}
	}

	return map[string]interface{}{key: v}
}
