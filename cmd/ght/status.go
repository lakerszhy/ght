package main

const (
	statusInProgress = iota
	statusSuccessful
	statusFailed
)

type status int

func (s status) IsInProgress() bool {
	return s == statusInProgress
}

func (s status) IsSuccessful() bool {
	return s == statusSuccessful
}

func (s status) IsFailed() bool {
	return s == statusFailed
}
