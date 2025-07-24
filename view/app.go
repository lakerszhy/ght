package view

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type app struct {
	panels []repoPanel
}

func NewApp() tea.Model {
	return app{
		panels: []repoPanel{
			newRepoPanel(),
			newRepoPanel(),
			newRepoPanel(),
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
	}

	for i := range a.panels {
		var cmd tea.Cmd
		a.panels[i], cmd = a.panels[i].Update(msg)
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
