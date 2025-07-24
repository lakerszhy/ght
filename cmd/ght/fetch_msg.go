package main

import (
	"net/http"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/lakerszhy/ght"
)

func FetchCmd() tea.Cmd {
	var cmds []tea.Cmd

	cmd := func() tea.Msg {
		return NewFetchInProgress()
	}
	cmds = append(cmds, cmd)

	cmd = func() tea.Msg {
		data, err := http.Get("https://github.com/trending/go?since=daily")
		if err != nil {
			return NewFetchFailed(err)
		}
		defer data.Body.Close()

		repos, err := ght.Parse(data.Body)
		if err != nil {
			return NewFetchFailed(err)
		}

		return NewFetchSuccessful(repos)
	}
	cmds = append(cmds, cmd)

	return tea.Sequence(cmds...)
}

type FetchMsg struct {
	Repos []ght.Repo
	Err   error
	status
}

func NewFetchInProgress() FetchMsg {
	return FetchMsg{
		status: statusInProgress,
	}
}

func NewFetchSuccessful(repos []ght.Repo) FetchMsg {
	return FetchMsg{
		Repos:  repos,
		status: statusSuccessful,
	}
}

func NewFetchFailed(err error) FetchMsg {
	return FetchMsg{
		Err:    err,
		status: statusFailed,
	}
}
