package handler

import (
	"github.com/shonjord/tracker/internal/pkg/domain/entity"
	"github.com/shonjord/tracker/internal/pkg/domain/query"
)

type (
	GetComputersAssignedToEmployee struct {
		finder employeeFinder
	}
)

// NewGetComputersAssignedToEmployee returns a new instance of this handler.
func NewGetComputersAssignedToEmployee(f employeeFinder) *GetComputersAssignedToEmployee {
	return &GetComputersAssignedToEmployee{
		finder: f,
	}
}

// GetComputersAssignedToEmployee returns a collection of computers of an employee.
func (h *GetComputersAssignedToEmployee) GetComputersAssignedToEmployee(
	q *query.GetComputersAssignedToEmployee,
) ([]*entity.Computer, error) {
	employee, err := h.finder.FindOneByUUID(q.EmployeeUUID)
	if err != nil {
		return nil, err
	}

	if employee.HasComputers() {
		return employee.Computers, nil
	}

	return []*entity.Computer{}, nil
}
