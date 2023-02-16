package action

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/shonjord/tracker/internal/pkg/domain/command"
	"github.com/shonjord/tracker/internal/pkg/presentation/rest"
)

type (
	computerToEmployeeUnassigner interface {
		Unassign(c *command.UnassignEmployeeFromComputer) error
	}

	UnassignNewComputerToEmployee struct {
		unassigner computerToEmployeeUnassigner
	}
)

// NewUnassignNewComputerToEmployee returns a new instance of this action.
func NewUnassignNewComputerToEmployee(
	u computerToEmployeeUnassigner,
) *UnassignNewComputerToEmployee {
	return &UnassignNewComputerToEmployee{
		unassigner: u,
	}
}

// Handle unassign a computer from an employee.
func (a *UnassignNewComputerToEmployee) Handle(req *rest.Request, res *rest.Response) error {
	var uuidParam string

	uuidParam = req.GetParam(computerUUIDParam)
	computerUUID, err := uuid.Parse(uuidParam)
	if err != nil {
		return res.BadRequest(&rest.Error{
			Message:               fmt.Sprintf("invalid computer UUID provided: %s", uuidParam),
			HTTPStatusCounterpart: http.StatusBadRequest,
		})
	}

	return a.unassigner.Unassign(&command.UnassignEmployeeFromComputer{
		ComputerUUID: computerUUID,
	})
}
