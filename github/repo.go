package github

import "fmt"

var (
	DateRangeDaily   = DateRange{Code: "daily", Name: "Today"}
	DateRangeWeekly  = DateRange{Code: "weekly", Name: "This week"}
	DateRangeMonthly = DateRange{Code: "monthly", Name: "This month"}
)

type DateRange struct {
	Code string
	Name string
}

type Repo struct {
	Owner         string
	Name          string
	Description   string
	Language      string
	LanguageColor string
	StarsTotal    string
	Forks         string
	StarsSince    string
}

func (r Repo) URL() string {
	return fmt.Sprintf("https://github.com/%s/%s", r.Owner, r.Name)
}
