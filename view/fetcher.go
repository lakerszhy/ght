package view

import (
	"net/http"
	"net/url"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/lakerszhy/ght/github"
)

func fetchCmd(language string, dateRange github.DateRange) tea.Cmd {
	var cmds []tea.Cmd

	cmd := func() tea.Msg {
		return newFetchInProgress(dateRange)
	}
	cmds = append(cmds, cmd)

	cmd = func() tea.Msg {
		u, err := url.Parse("https://github.com/trending")
		if err != nil {
			return newFetchFailed(dateRange, err)
		}

		if language != "" {
			u = u.JoinPath(language)
		}

		params := url.Values{}
		params.Set("since", dateRange.Code)

		u.RawQuery = params.Encode()

		data, err := http.Get(u.String())
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
