package docker

import (
	"github.com/shirou/gopsutil/docker"
)

type InfoDocker struct {
	Docker []docker.CgroupDockerStat `json:"docker"`
}

func (i *InfoDocker) Run() (value map[string]interface{}) {
	key := "docker"
	v, err := docker.GetDockerStat()
	if err != nil {
		return map[string]interface{}{key: err}
	}

	return map[string]interface{}{key: v}
}
