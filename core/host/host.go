package host

import (
	"github.com/shirou/gopsutil/host"
)

type InfoHost struct {
	Host *host.InfoStat `json:"host"`
}

func (i *InfoHost) Run() (value map[string]interface{}) {
	key := "host"
	v, err := host.Info()
	if err != nil {
		return map[string]interface{}{key: err}
	}

	return map[string]interface{}{key: v}
}
