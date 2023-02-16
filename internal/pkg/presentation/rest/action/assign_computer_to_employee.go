package action

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/shonjord/tracker/internal/pkg/domain/command"
	"github.com/shonjord/tracker/internal/pkg/presentation/rest"
)

type (
	computerToEmployeeAssigner interface {
		Assign(*command.AssignComputerToEmployee) error
	}
	AssignComputerToEmployee struct {
		assigner computerToEmployeeAssigner
	}
)

// NewAssignComputerToEmployee returns a new instance of this action.
func NewAssignComputerToEmployee(a computerToEmployeeAssigner) *AssignComputerToEmployee {
	return &AssignComputerToEmployee{
		assigner: a,
	}
}

// Handle adds a new computer to an employee.
func (a *AssignComputerToEmployee) Handle(req *rest.Request, res *rest.Response) error {
	var cmd *command.AssignComputerToEmployee

	uuidParam := req.GetParam(computerUUIDParam)
	computerUUID, err := uuid.Parse(uuidParam)
	if err != nil {
		return res.BadRequest(&rest.Error{
			Message:               fmt.Sprintf("invalid computer UUID provided: %s", uuidParam),
			HTTPStatusCounterpart: http.StatusBadRequest,
		})
	}

	if err = json.NewDecoder(req.Body()).Decode(&cmd); err != nil {
		return res.BadRequest(err)
	}

	cmd.ComputerUUID = computerUUID

	return a.assigner.Assign(cmd)
}
