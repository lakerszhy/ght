package view

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lakerszhy/ght/github"
)

type app struct {
	panels []repoPanel
	focus  github.DateRange
}

func NewApp() tea.Model {
	return app{
		focus: github.DateRangeDaily,
		panels: []repoPanel{
			newRepoPanel(github.DateRangeDaily, true),
			newRepoPanel(github.DateRangeWeekly, false),
			newRepoPanel(github.DateRangeMonthly, false),
		},
	}
}

func (a app) Init() tea.Cmd {
	var cmds []tea.Cmd
	for _, panel := range a.panels {
		cmds = append(cmds, panel.Init())
	}
	return tea.Batch(cmds...)
}

func (a app) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case fetchMsg:
		for i := range a.panels {
			var cmd tea.Cmd
			a.panels[i], cmd = a.panels[i].Update(msg)
			cmds = append(cmds, cmd)
		}
		return a, tea.Batch(cmds...)
	case tea.WindowSizeMsg:
		panels := make([]repoPanel, 0, len(a.panels))
		for _, p := range a.panels {
			p = p.SetSize(msg.Width/len(a.panels), msg.Height)
			panels = append(panels, p)
		}
		a.panels = panels
		return a, nil
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return a, tea.Quit
		}
		if msg.String() == "tab" {
			switch a.focus {
			case github.DateRangeDaily:
				a.focus = github.DateRangeWeekly
			case github.DateRangeWeekly:
				a.focus = github.DateRangeMonthly
			case github.DateRangeMonthly:
				a.focus = github.DateRangeDaily
			}
			for i := range a.panels {
				a.panels[i].IsFocused = a.panels[i].dateRange == a.focus
			}
			return a, nil
		}
	}

	for i := range a.panels {
		var cmd tea.Cmd
		if a.panels[i].IsFocused {
			a.panels[i], cmd = a.panels[i].Update(msg)
		}
		cmds = append(cmds, cmd)
	}
	return a, tea.Batch(cmds...)
}

func (a app) View() string {
	var views []string
	for _, l := range a.panels {
		views = append(views, l.View())
	}
	return lipgloss.JoinHorizontal(lipgloss.Left, views...)
}
