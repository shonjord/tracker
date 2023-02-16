package rest

type (
	Error struct {
		Message               string
		HTTPStatusCounterpart int
	}
)

// Error returns a string representation of this error.
func (e *Error) Error() string {
	return e.Message
}
