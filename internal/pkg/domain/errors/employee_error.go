package errors

import (
	"fmt"

	"github.com/google/uuid"
)

// NewEmployeeForUUIDNotFound returns an error for when an employee for the given UUID is not found.
func NewEmployeeForUUIDNotFound(uuid uuid.UUID) *NotFoundError {
	return &NotFoundError{
		Message: fmt.Sprintf("employee with UUID: %s, not found", uuid),
	}
}
