package view

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/lakerszhy/ght/github"
)

func fetchCmd() tea.Cmd {
	var cmds []tea.Cmd

	cmd := func() tea.Msg {
		return newFetchInProgress()
	}
	cmds = append(cmds, cmd)

	cmd = func() tea.Msg {
		// data, err := http.Get("https://github.com/trending/go?since=daily")
		// if err != nil {
		// 	return newFetchFailed(err)
		// }
		// defer data.Body.Close()

		// TODO: mock
		f, err := os.Open("github/testdata/go_daily.html")
		if err != nil {
			return newFetchFailed(err)
		}
		defer f.Close()

		repos, err := github.Parse(f)
		if err != nil {
			return newFetchFailed(err)
		}

		return newFetchSuccessful(repos)
	}
	cmds = append(cmds, cmd)

	return tea.Sequence(cmds...)
}

type fetchMsg struct {
	Repos []github.Repo
	Err   error
	status
}

func newFetchInProgress() fetchMsg {
	return fetchMsg{
		status: statusInProgress,
	}
}

func newFetchSuccessful(repos []github.Repo) fetchMsg {
	return fetchMsg{
		Repos:  repos,
		status: statusSuccessful,
	}
}

func newFetchFailed(err error) fetchMsg {
	return fetchMsg{
		Err:    err,
		status: statusFailed,
	}
}
