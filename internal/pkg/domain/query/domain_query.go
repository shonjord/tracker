package query

import "github.com/google/uuid"

type (
	GetAllComputers struct {
		Limit int
	}

	GetComputerByUUID struct {
		ComputerUUID uuid.UUID
	}

	GetComputersAssignedToEmployee struct {
		EmployeeUUID uuid.UUID
	}
)
