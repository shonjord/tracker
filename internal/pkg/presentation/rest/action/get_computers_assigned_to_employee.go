package action

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/shonjord/tracker/internal/pkg/domain/entity"
	"github.com/shonjord/tracker/internal/pkg/domain/query"
	"github.com/shonjord/tracker/internal/pkg/presentation/rest"
)

type (
	computersAssignedToEmployeeGetter interface {
		GetComputersAssignedToEmployee(*query.GetComputersAssignedToEmployee) ([]*entity.Computer, error)
	}
	GetComputersAssignedToEmployee struct {
		getter computersAssignedToEmployeeGetter
	}
)

// NewGetComputersAssignedToEmployee returns a new instance of this action.
func NewGetComputersAssignedToEmployee(
	g computersAssignedToEmployeeGetter,
) *GetComputersAssignedToEmployee {
	return &GetComputersAssignedToEmployee{
		getter: g,
	}
}

// Handle returns all computers linked to an employee.
func (a *GetComputersAssignedToEmployee) Handle(req *rest.Request, res *rest.Response) error {
	var (
		cmd = new(query.GetComputersAssignedToEmployee)
	)

	uuidParam := req.GetParam(employeeUUIDParam)
	employeeUUID, err := uuid.Parse(uuidParam)
	if err != nil {
		return res.BadRequest(&rest.Error{
			Message:               fmt.Sprintf("invalid UUID provided: %s", uuidParam),
			HTTPStatusCounterpart: http.StatusBadRequest,
		})
	}

	cmd.EmployeeUUID = employeeUUID

	computers, err := a.getter.GetComputersAssignedToEmployee(cmd)
	if err != nil {
		return err
	}

	return res.WriteStruct(computers)
}
