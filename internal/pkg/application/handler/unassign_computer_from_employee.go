package handler

import (
	"github.com/shonjord/tracker/internal/pkg/domain/command"
	"github.com/shonjord/tracker/internal/pkg/domain/errors"
)

type (
	UnassignComputerFromEmployee struct {
		finder        employeeFinder
		finderUpdater computerFinderUpdater
	}
)

// NewUnassignComputerFromEmployee returns a new instance of this handler.
func NewUnassignComputerFromEmployee(
	f employeeFinder,
	fu computerFinderUpdater,
) *UnassignComputerFromEmployee {
	return &UnassignComputerFromEmployee{
		finder:        f,
		finderUpdater: fu,
	}
}

// Unassign removes the current employee from the computer.
func (h *UnassignComputerFromEmployee) Unassign(c *command.UnassignEmployeeFromComputer) error {
	computer, err := h.finderUpdater.FindOneByUUID(c.ComputerUUID)
	if err != nil {
		return err
	}
	if !computer.HasEmployee() {
		return errors.NewComputerHasNoEmployeeAssigned(computer)
	}

	computer.UnassignEmployee()

	return h.finderUpdater.Update(computer)
}
