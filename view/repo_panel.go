package view

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/ansi"
	"github.com/lakerszhy/ght/github"
	"github.com/pkg/browser"
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
		fetchMsg:  newFetchInProgress(dateRange),
		list:      list,
		IsFocused: isFocused,
	}
}

func (p repoPanel) Init(language string) tea.Cmd {
	return fetchCmd(language, p.dateRange)
}

func (p repoPanel) Update(msg tea.Msg) (repoPanel, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case fetchMsg:
		if msg.DateRange != p.dateRange {
			return p, nil
		}
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
	case tea.KeyMsg:
		if msg.String() == "o" || msg.String() == "enter" {
			if i, ok := p.list.SelectedItem().(repoItem); ok {
				browser.OpenURL(i.URL())
			}
			return p, nil
		}
	}

	p.list, cmd = p.list.Update(msg)
	return p, cmd
}

func (p repoPanel) View() string {
	borderColor := "#3d444d"
	if p.IsFocused {
		borderColor = "#4493f8"
	}
	style := lipgloss.NewStyle().Border(p.border()).
		BorderForeground(lipgloss.Color(borderColor)).
		Width(p.list.Width()).Height(p.list.Height())

	if p.fetchMsg.isInProgress() {
		return style.AlignHorizontal(lipgloss.Center).Render("Fetching...")
	}

	if p.fetchMsg.isFailed() {
		return style.AlignHorizontal(lipgloss.Center).
			Render(fmt.Sprintf("Failed: %s", p.fetchMsg.Err))
	}

	return style.Render(p.list.View())
}

func (p repoPanel) border() lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border = p.borderTitle(border)
	border = p.borderBottom(border)
	return border
}

func (p repoPanel) borderTitle(border lipgloss.Border) lipgloss.Border {
	title := fmt.Sprintf("%s %s %s", border.MiddleRight, p.dateRange.Name, border.MiddleLeft)
	repeatCount := max(p.list.Width()-ansi.StringWidth(title)-ansi.StringWidth(border.Top), 0)
	end := strings.Repeat(border.Top, repeatCount)
	top := border.Top + title + end
	border.Top = top
	return border
}

func (p repoPanel) borderBottom(border lipgloss.Border) lipgloss.Border {
	if len(p.list.Items()) == 0 {
		return border
	}

	foot := fmt.Sprintf("%d/%d", p.list.Index()+1, len(p.list.Items()))
	foot = fmt.Sprintf("%s%s%s", border.MiddleRight, foot, border.MiddleLeft)
	repeatCount := max(p.list.Width()-ansi.StringWidth(foot)-ansi.StringWidth(border.Bottom), 0)
	start := strings.Repeat(border.Bottom, repeatCount)
	bottom := start + foot + border.Bottom
	border.Bottom = bottom
	return border
}

func (p repoPanel) SetSize(width int, height int) repoPanel {
	p.list.SetSize(width-2, height-2)
	return p
}
