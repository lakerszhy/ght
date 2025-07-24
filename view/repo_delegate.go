package view

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/lakerszhy/ght/github"
)

type repoItem struct {
	github.Repo
}

func (r repoItem) FilterValue() string {
	return r.Name
}

type repoDelegate struct {
}

func NewRepoDelegate() list.ItemDelegate {
	return repoDelegate{}
}

func (d repoDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	i, ok := item.(repoItem)
	if !ok {
		return
	}
	fmt.Fprintf(w, "%s/%s", i.Owner, i.Name)
}

func (d repoDelegate) Height() int {
	return 1
}

func (d repoDelegate) Spacing() int {
	return 0
}

func (d repoDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd {
	return nil
}
