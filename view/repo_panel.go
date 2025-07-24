package view

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lakerszhy/ght/github"
)

type repoPanel struct {
	dateRange github.DateRange
	fetchMsg  fetchMsg
	list      list.Model
	IsFocused bool
}

func newRepoPanel(dateRange github.DateRange, isFocused bool) repoPanel {
	list := list.New([]list.Item{}, NewRepoDelegate(), 0, 0)
	list.DisableQuitKeybindings()
	list.SetShowTitle(false)
	list.SetFilteringEnabled(false)
	list.SetShowPagination(false)
	list.SetShowStatusBar(false)
	list.SetShowHelp(false)
	return repoPanel{
		dateRange: dateRange,
		fetchMsg:  newFetchInProgress(),
		list:      list,
		IsFocused: isFocused,
	}
}

func (p repoPanel) Init() tea.Cmd {
	return fetchCmd()
}

func (p repoPanel) Update(msg tea.Msg) (repoPanel, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case fetchMsg:
		p.fetchMsg = msg
		if p.fetchMsg.isSuccessful() {
			repos := p.fetchMsg.Repos
			items := make([]list.Item, 0, len(repos))
			for _, repo := range repos {
				items = append(items, repoItem{repo})
			}
			cmd = p.list.SetItems(items)
		}
		return p, cmd
	}

	p.list, cmd = p.list.Update(msg)
	return p, cmd
}

func (p repoPanel) View() string {
	if p.fetchMsg.isInProgress() {
		return "Fetching..."
	}

	if p.fetchMsg.isFailed() {
		return fmt.Sprintf("Failed: %s", p.fetchMsg.Err)
	}

	borderColor := "#3d444d"
	if p.IsFocused {
		borderColor = "#4493f8"
	}
	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(borderColor)).
		Width(p.list.Width()).Render(p.list.View())
}

func (p repoPanel) SetSize(width int, height int) repoPanel {
	p.list.SetSize(width-2, height-2)
	return p
}
