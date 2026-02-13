package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"

	"github.com/yunnij/sysbeat/internal/tui"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run interactive TUI monitor",
	RunE: func(cmd *cobra.Command, args []string) error {
		p := tea.NewProgram(tui.NewModel())
		_, err := p.Run()
		return err
	},
}
