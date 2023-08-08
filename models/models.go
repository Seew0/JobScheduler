package models

type State int

const (
	Idle State = iota
	Running
	Done
)
type Job struct {
	Duration  int
	Status    State
	StartTime int64
}
