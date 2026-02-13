package tui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/yunnij/sysbeat/internal/collect"
)

type tickMsg time.Time
type metricsMsg struct {
	m   collect.Metrics
	err error
}

type Model struct {
	metrics collect.Metrics
	err     error
}

func NewModel() Model {
	return Model{}
}

func (m Model) Init() tea.Cmd {
	// 시작하자마자 1초 타이머 + 최초 수집 실행
	return tea.Batch(tick(), fetchMetrics())
}

func tick() tea.Cmd {
	return tea.Tick(1*time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func fetchMetrics() tea.Cmd {
	return func() tea.Msg {
		mt, err := collect.CollectOnce()
		return metricsMsg{m: mt, err: err}
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tickMsg:
		// 1초마다: 다음 tick 예약 + 메트릭 다시 수집
		return m, tea.Batch(tick(), fetchMetrics())

	case metricsMsg:
		m.metrics = msg.m
		m.err = msg.err
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, nil
}
