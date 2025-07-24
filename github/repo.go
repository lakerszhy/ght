package github

const (
	DateRangeDaily   DateRange = "daily"
	DateRangeWeekly  DateRange = "weekly"
	DateRangeMonthly DateRange = "monthly"
)

type DateRange string

func (d DateRange) String() string {
	switch d {
	case DateRangeDaily:
		return "Today"
	case DateRangeWeekly:
		return "This week"
	case DateRangeMonthly:
		return "This month"
	default:
		return ""
	}
}

type Repo struct {
	Owner       string
	Name        string
	Description string
	Language    string
	StarsTotal  string
	Forks       string
	StarsSince  string
}
