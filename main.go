package main

import (
	"flag"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/lakerszhy/ght/view"
)

var version = "dev"

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	var language string
	flag.StringVar(&language, "l", "", "filter repos by language")
	flag.Parse()

	p := tea.NewProgram(view.NewApp(language, version),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)
	if _, err := p.Run(); err != nil {
		return err
	}

	return nil
}
