package tui

import "fmt"

func (m Model) View() string {
	if m.err != nil {
		return fmt.Sprintf("SysBeat\n\nerror: %v\n\n(press q to quit)\n", m.err)
	}

	return fmt.Sprintf(
		"SysBeat (press q to quit)\n\nCPU:  %.1f%%\nMEM:  %.1f%%\nDISK: %.1f%%\n",
		m.metrics.CPUPercent,
		m.metrics.MemUsedPercent,
		m.metrics.DiskUsedPercent,
	)
}
