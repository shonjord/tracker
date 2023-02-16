package command

import "github.com/google/uuid"

type (
	AssignNewComputerToEmployee struct {
		MacAddress   string
		IPAddress    string
		Name         string
		Description  string
		EmployeeUUID uuid.UUID
	}

	UnassignEmployeeFromComputer struct {
		ComputerUUID uuid.UUID
	}

	AssignComputerToEmployee struct {
		ComputerUUID uuid.UUID
		EmployeeUUID uuid.UUID
	}
)
