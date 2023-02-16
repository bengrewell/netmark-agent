package process

import (
	"github.com/shirou/gopsutil/process"
)

type ProcessInfo struct {
	Pid        int32                   `json:"pid"`
	Ppid       int32                   `json:"ppid"`
	Name       string                  `json:"name"`
	Cmdline    string                  `json:"cmdline"`
	CreateTime int64                   `json:"created_time"`
	Exe        string                  `json:"exe"`
	IoCounters *process.IOCountersStat `json:"io_counters"`
	Nice       int32                   `json:"nice"`
	NumThreads int32                   `json:"num_threads"`
	MemoryInfo *process.MemoryInfoStat `json:"memory_info"`
	Username   string                  `json:"username"`
}

func (pi *ProcessInfo) Update(p *process.Process) (err error) {
	pi.Pid = p.Pid
	pi.Ppid, _ = p.Ppid()
	pi.Name, _ = p.Name()
	pi.Cmdline, _ = p.Cmdline()
	pi.CreateTime, _ = p.CreateTime()
	pi.Exe, _ = p.Exe()
	pi.IoCounters, _ = p.IOCounters()
	pi.Nice, _ = p.Nice()
	pi.NumThreads, _ = p.NumThreads()
	pi.MemoryInfo, _ = p.MemoryInfo()
	pi.Username, _ = p.Username()
	return nil
}

type InfoProcess struct {
	Processes []*ProcessInfo `json:"processes"`
}

func (i *InfoProcess) Run() (value map[string]interface{}) {
	key := "process"
	i.Processes = make([]*ProcessInfo, 0)
	processes, err := process.Processes()
	if err != nil {
		i.Processes = nil
	} else {
		for _, p := range processes {
			pi := ProcessInfo{}
			pi.Update(p)
			i.Processes = append(i.Processes, &pi)
		}
	}

	return map[string]interface{}{key: i.Processes}
}
