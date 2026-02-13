package app

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/yunnij/sysbeat/internal/collect"
)

var snapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "Print a single metrics snapshot",
	RunE: func(cmd *cobra.Command, args []string) error {
		m, err := collect.CollectOnce()
		if err != nil {
			return err
		}

		fmt.Printf(
			"CPU: %.1f%% | MEM: %.1f%% | DISK: %.1f%%\n",
			m.CPUPercent,
			m.MemUsedPercent,
			m.DiskUsedPercent,
		)
		return nil
	},
}
