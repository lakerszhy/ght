package view

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type repoPanel struct {
	fetchMsg fetchMsg
	list     list.Model
}

func newRepoPanel() repoPanel {
	list := list.New([]list.Item{}, NewRepoDelegate(), 0, 0)
	list.DisableQuitKeybindings()
	return repoPanel{
		fetchMsg: newFetchInProgress(),
		list:     list,
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

	return p.list.View()
}

func (p repoPanel) SetSize(width int, height int) repoPanel {
	p.list.SetSize(width, height)
	return p
}
