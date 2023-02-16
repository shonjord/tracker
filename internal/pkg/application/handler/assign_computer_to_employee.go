package handler

import (
	"github.com/shonjord/tracker/internal/pkg/domain/command"
	"github.com/shonjord/tracker/internal/pkg/domain/errors"
	log "github.com/sirupsen/logrus"
)

type (
	AssignComputerToEmployee struct {
		finder        employeeFinder
		finderUpdater computerFinderUpdater
		notifier      adminNotifier
	}
)

// NewAssignComputerToEmployee returns a new instance of this handler.
func NewAssignComputerToEmployee(
	f employeeFinder,
	fs computerFinderUpdater,
	n adminNotifier,
) *AssignComputerToEmployee {
	return &AssignComputerToEmployee{
		finder:        f,
		finderUpdater: fs,
		notifier:      n,
	}
}

// Assign assigns an existing computer to a different employee.
func (h *AssignComputerToEmployee) Assign(c *command.AssignComputerToEmployee) error {
	computer, err := h.finderUpdater.FindOneByUUID(c.ComputerUUID)
	if err != nil {
		return err
	}
	if computer.HasEmployee() {
		return errors.NewComputerHasEmployeeAssigned(computer)
	}

	employee, err := h.finder.FindOneByUUID(c.EmployeeUUID)
	if err != nil {
		return err
	}

	computer.WithEmployee(employee)

	if err = h.finderUpdater.Update(computer); err != nil {
		return err
	}

	if employee.HasBeenAssignedWithThreeOrMoreComputers() {
		if err = h.notifier.Notify(computer); err != nil {
			log.WithError(err).Error("error while notifying to admin service")
		}
	}

	return nil
}
