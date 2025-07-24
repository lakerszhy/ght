package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/lakerszhy/ght/view"
)

func main() {
	p := tea.NewProgram(view.NewApp(),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
