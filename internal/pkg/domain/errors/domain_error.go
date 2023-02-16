package errors

type (
	ConflictError struct {
		Message string
	}
	NotFoundError struct {
		Message string
	}
	InternalError struct {
		Message string
	}
)

// Error returns a string representation of a conflict domain error.
func (e *ConflictError) Error() string {
	return e.Message
}

// Error returns a string representation of a not found domain error.
func (e *NotFoundError) Error() string {
	return e.Message
}

// Error returns a string representation of an internal error.
func (e *InternalError) Error() string {
	return e.Message
}
