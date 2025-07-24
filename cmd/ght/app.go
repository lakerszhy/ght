package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/lakerszhy/ght"
)

type app struct {
	fetchMsg FetchMsg
	list     list.Model
}

func newApp() app {
	list := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	list.DisableQuitKeybindings()
	return app{
		fetchMsg: NewFetchInProgress(),
		list:     list,
	}
}

func (a app) Init() tea.Cmd {
	return FetchCmd()
}

func (a app) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case FetchMsg:
		a.fetchMsg = msg
		if a.fetchMsg.IsSuccessful() {
			repos := a.fetchMsg.Repos
			items := make([]list.Item, 0, len(repos))
			for _, repo := range repos {
				items = append(items, repoItem{repo: repo})
			}
			a.list.SetItems(items)
		}
		return a, nil
	case tea.WindowSizeMsg:
		a.list.SetSize(msg.Width, msg.Height)
		return a, nil
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return a, tea.Quit
		}
	}

	var cmd tea.Cmd
	a.list, cmd = a.list.Update(msg)
	return a, cmd
}

func (a app) View() string {
	if a.fetchMsg.IsInProgress() {
		return "Fetching..."
	}

	if a.fetchMsg.IsFailed() {
		return fmt.Sprintf("Failed: %s", a.fetchMsg.Err)
	}

	return a.list.View()
}

type repoItem struct {
	repo ght.Repo
}

func (r repoItem) Title() string {
	return fmt.Sprintf("%s/%s", r.repo.Owner, r.repo.Name)
}

func (r repoItem) Description() string {
	return r.repo.Description
}

func (r repoItem) FilterValue() string {
	return r.repo.Name
}
