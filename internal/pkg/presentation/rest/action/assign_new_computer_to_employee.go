package action

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/shonjord/tracker/internal/pkg/domain/command"
	"github.com/shonjord/tracker/internal/pkg/domain/entity"
	"github.com/shonjord/tracker/internal/pkg/presentation/rest"
)

type (
	newComputerToEmployeeAssigner interface {
		Assign(*command.AssignNewComputerToEmployee) (*entity.Computer, error)
	}
	AssignNewComputerToEmployee struct {
		assigner newComputerToEmployeeAssigner
	}
)

// NewAssignNewComputerToEmployee returns a new instance of this action.
func NewAssignNewComputerToEmployee(a newComputerToEmployeeAssigner) *AssignNewComputerToEmployee {
	return &AssignNewComputerToEmployee{
		assigner: a,
	}
}

// Handle adds a new computer to an employee.
func (a *AssignNewComputerToEmployee) Handle(req *rest.Request, res *rest.Response) error {
	var cmd *command.AssignNewComputerToEmployee

	uuidParam := req.GetParam(employeeUUIDParam)
	employeeUUID, err := uuid.Parse(uuidParam)
	if err != nil {
		return res.BadRequest(&rest.Error{
			Message:               fmt.Sprintf("invalid UUID provided: %s", uuidParam),
			HTTPStatusCounterpart: http.StatusBadRequest,
		})
	}

	if err = json.NewDecoder(req.Body()).Decode(&cmd); err != nil {
		return res.BadRequest(err)
	}

	cmd.EmployeeUUID = employeeUUID

	computer, err := a.assigner.Assign(cmd)
	if err != nil {
		return err
	}

	return res.WriteStruct(computer)
}
