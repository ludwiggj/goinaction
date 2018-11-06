package counters

// unexported type
type alertCounter int
type AlertCounter int

// New creates and returns values of the unexported type alertCounter
// It’s a convention in Go to give factory functions the name of New
func New(value int) alertCounter {
	return alertCounter(value)
}
