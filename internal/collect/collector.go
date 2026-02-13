package collect

import (
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

func CollectOnce() (Metrics, error) {
	var out Metrics

	// CPU는 짧은 시간 샘플링이 필요해서 200ms 측정
	cpuPercents, err := cpu.Percent(200*time.Millisecond, false)
	if err != nil {
		return out, err
	}
	if len(cpuPercents) > 0 {
		out.CPUPercent = cpuPercents[0]
	}

	vm, err := mem.VirtualMemory()
	if err != nil {
		return out, err
	}
	out.MemUsedPercent = vm.UsedPercent

	du, err := disk.Usage("/")
	if err != nil {
		return out, err
	}
	out.DiskUsedPercent = du.UsedPercent

	return out, nil
}
