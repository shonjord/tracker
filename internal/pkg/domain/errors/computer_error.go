package errors

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/shonjord/tracker/internal/pkg/domain/entity"
)

// NewMacAddressIsAssignedToAnotherEmployee means a conflict error.
func NewMacAddressIsAssignedToAnotherEmployee(c *entity.Computer) *ConflictError {
	return &ConflictError{
		Message: fmt.Sprintf(
			"mac: %s, has been assigned already, can't be assigned to employe: %s",
			c.MACAddress,
			c.Employee.UUID,
		),
	}
}

// NewComputerHasNoEmployeeAssigned means a conflict error.
func NewComputerHasNoEmployeeAssigned(c *entity.Computer) *ConflictError {
	return &ConflictError{
		Message: fmt.Sprintf(
			"computer with UUID: %s, has no employee assigned",
			c.UUID,
		),
	}
}

// NewComputerHasEmployeeAssigned means a conflict error.
func NewComputerHasEmployeeAssigned(c *entity.Computer) *ConflictError {
	return &ConflictError{
		Message: fmt.Sprintf(
			"computer with UUID: %s, has an employee already assigned",
			c.UUID,
		),
	}
}

// NewComputerForUUIDNotFound returns an error for when a computer for the given UUID is not found.
func NewComputerForUUIDNotFound(uuid uuid.UUID) *NotFoundError {
	return &NotFoundError{
		Message: fmt.Sprintf("computer with UUID: %s, not found", uuid),
	}
}
