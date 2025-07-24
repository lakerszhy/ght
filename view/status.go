package view

const (
	statusInProgress = iota
	statusSuccessful
	statusFailed
)

type status int

func (s status) isInProgress() bool {
	return s == statusInProgress
}

func (s status) isSuccessful() bool {
	return s == statusSuccessful
}

func (s status) isFailed() bool {
	return s == statusFailed
}
