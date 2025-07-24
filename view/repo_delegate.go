package view

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/ansi"
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

	title := d.title(i)
	title = ansi.Truncate(title, m.Width()-3, "...")

	desc := d.description(i)
	desc = ansi.Truncate(desc, m.Width()-3, "...")

	info := d.info(i)

	style := lipgloss.NewStyle().Padding(0, 1, 0, 2)
	if m.Index() == index {
		style = style.Border(lipgloss.RoundedBorder(),
			false, false, false, true).
			BorderForeground(lipgloss.Color("#4493f8")).
			Padding(0, 1)
	}

	v := style.Render(lipgloss.JoinVertical(
		lipgloss.Left, title, desc, info))
	fmt.Fprintf(w, "%s", v)
}

func (d repoDelegate) title(item repoItem) string {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("#4493f8"))
	owner := style.Render(item.Owner)
	name := style.Bold(true).Render(fmt.Sprintf("/%s", item.Name))

	return fmt.Sprintf("%s%s", owner, name)
}

func (d repoDelegate) description(item repoItem) string {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("#9198a1"))
	return style.Render(item.Description)
}

func (d repoDelegate) info(item repoItem) string {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("#9198a1"))

	language := style.Render(item.Language)
	stars := style.Render(item.StarsTotal)
	forks := style.Render(item.Forks)

	return fmt.Sprintf("%s %s %s", language, stars, forks)
}

func (d repoDelegate) Height() int {
	return 3
}

func (d repoDelegate) Spacing() int {
	return 1
}

func (d repoDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd {
	return nil
}
