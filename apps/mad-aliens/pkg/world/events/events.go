package events

type Event string

const (
	Loaded    Event = "loaded"
	Moved     Event = "moved"
	CleanedUp Event = "cleaned_up"
)
