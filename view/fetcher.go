package view

import (
	"fmt"
	"net/http"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/lakerszhy/ght/github"
)

func fetchCmd(dateRange github.DateRange) tea.Cmd {
	var cmds []tea.Cmd

	cmd := func() tea.Msg {
		return newFetchInProgress(dateRange)
	}
	cmds = append(cmds, cmd)

	cmd = func() tea.Msg {
		url := fmt.Sprintf("https://github.com/trending/?since=%s", dateRange.Code)
		data, err := http.Get(url)
		if err != nil {
			return newFetchFailed(dateRange, err)
		}
		defer data.Body.Close()

		// f, err := os.Open("github/testdata/go_daily.html")
		// if err != nil {
		// 	return newFetchFailed(dateRange, err)
		// }
		// defer f.Close()

		repos, err := github.Parse(data.Body)
		if err != nil {
			return newFetchFailed(dateRange, err)
		}

		return newFetchSuccessful(dateRange, repos)
	}
	cmds = append(cmds, cmd)

	return tea.Sequence(cmds...)
}

type fetchMsg struct {
	DateRange github.DateRange
	Repos     []github.Repo
	Err       error
	status
}

func newFetchInProgress(dateRange github.DateRange) fetchMsg {
	return fetchMsg{
		DateRange: dateRange,
		status:    statusInProgress,
	}
}

func newFetchSuccessful(dateRange github.DateRange, repos []github.Repo) fetchMsg {
	return fetchMsg{
		DateRange: dateRange,
		Repos:     repos,
		status:    statusSuccessful,
	}
}

func newFetchFailed(dateRange github.DateRange, err error) fetchMsg {
	return fetchMsg{
		DateRange: dateRange,
		Err:       err,
		status:    statusFailed,
	}
}
