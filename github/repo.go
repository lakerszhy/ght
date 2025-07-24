package github

type DateRange string

const (
	DateRangeDaily   DateRange = "daily"
	DateRangeWeekly  DateRange = "weekly"
	DateRangeMonthly DateRange = "monthly"
)

type Repo struct {
	Owner       string
	Name        string
	Description string
	Language    string
	StarsTotal  string
	Forks       string
	StarsSince  string
}
