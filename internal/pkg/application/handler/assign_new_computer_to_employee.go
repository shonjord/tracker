package handler

import (
	"github.com/shonjord/tracker/internal/pkg/domain/command"
	"github.com/shonjord/tracker/internal/pkg/domain/entity"
	log "github.com/sirupsen/logrus"
)

type (
	AssignNewComputerToEmployee struct {
		finder   employeeFinder
		saver    computerSaver
		notifier adminNotifier
	}
)

// NewAssignNewComputerToEmployee returns a new instance of this handler.
func NewAssignNewComputerToEmployee(
	f employeeFinder,
	s computerSaver,
	n adminNotifier,
) *AssignNewComputerToEmployee {
	return &AssignNewComputerToEmployee{
		finder:   f,
		saver:    s,
		notifier: n,
	}
}

// Assign adds new computer to an employee.
func (h *AssignNewComputerToEmployee) Assign(
	c *command.AssignNewComputerToEmployee,
) (*entity.Computer, error) {

	employee, err := h.finder.FindOneByUUID(c.EmployeeUUID)
	if err != nil {
		return nil, err
	}

	computer := entity.NewComputer(
		c.MacAddress,
		c.IPAddress,
		c.Name,
		c.Description,
		employee,
	)

	employee.AddComputer(computer)

	if err = h.saver.Save(computer); err != nil {
		return nil, err
	}

	if employee.HasBeenAssignedWithThreeOrMoreComputers() {
		if err = h.notifier.Notify(computer); err != nil {
			log.WithError(err).Error("error while notifying to admin service")
		}
	}

	return computer, nil
}
