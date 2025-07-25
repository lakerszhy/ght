package view

import (
	"fmt"
	"io"
	"strings"

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

	title := ansi.Truncate(d.title(i), m.Width()-3, "...")
	desc := ansi.Truncate(d.description(i), m.Width()-3, "...")
	info := ansi.Truncate(d.info(i), m.Width()-3, "...")

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

	infos := []string{}

	if item.Language != "" {
		dot := lipgloss.NewStyle().Foreground(lipgloss.Color(item.LanguageColor)).Render("●")
		language := style.Render(item.Language)
		infos = append(infos, fmt.Sprintf("%s %s", dot, language))
	}

	infos = append(infos, fmt.Sprintf("⭑ %s", item.StarsTotal))
	infos = append(infos, fmt.Sprintf("🝘 %s", item.Forks))

	return strings.Join(infos, "   ")
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
