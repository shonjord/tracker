package handler_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/shonjord/tracker/internal/pkg/application/handler"
	"github.com/shonjord/tracker/internal/pkg/domain/command"
	"github.com/shonjord/tracker/internal/pkg/domain/entity"
	"github.com/shonjord/tracker/internal/pkg/domain/errors"
	assert "github.com/shonjord/tracker/test/.assert"
	mock "github.com/shonjord/tracker/test/.mock"
)

func TestUnassignComputerFromEmployee(t *testing.T) {
	t.Parallel()

	tests := []struct {
		scenario string
		function func(
			*testing.T,
			*mock.ComputerRepository,
			*mock.EmployeeRepository,
			*command.UnassignEmployeeFromComputer,
		)
	}{
		{
			// Given a valid domain command
			scenario: "When computer repository can't find a computer",
			function: thenUnassignComputerFromEmployeeShouldReturnComputerNotFoundError,
		},
		{
			// Given a valid domain command
			scenario: "When computer repository finds a computer but has no employee assigned",
			function: thenUnassignComputerFromEmployeeShouldReturnConflictError,
		},
		{
			// Given a valid domain command
			scenario: "When computer repository finds a computer that has an employee assigned",
			function: thenUnassignComputerFromEmployeeShouldUnassignEmployeeSuccessfully,
		},
	}

	cmd := &command.UnassignEmployeeFromComputer{
		ComputerUUID: uuid.New(),
	}

	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			test.function(
				t,
				new(mock.ComputerRepository),
				new(mock.EmployeeRepository),
				cmd,
			)
		})
	}
}

func thenUnassignComputerFromEmployeeShouldReturnComputerNotFoundError(
	t *testing.T,
	cr *mock.ComputerRepository,
	er *mock.EmployeeRepository,
	cmd *command.UnassignEmployeeFromComputer,
) {
	cr.FindOneByUUIDFunc = func(uuid uuid.UUID) (*entity.Computer, error) {
		return nil, errors.NewComputerForUUIDNotFound(uuid)
	}

	h := handler.NewUnassignComputerFromEmployee(er, cr)
	err := h.Unassign(cmd)

	assert.True(t, err != nil, "should fail with computer repository errors.")
}

func thenUnassignComputerFromEmployeeShouldReturnConflictError(
	t *testing.T,
	cr *mock.ComputerRepository,
	er *mock.EmployeeRepository,
	cmd *command.UnassignEmployeeFromComputer,
) {
	cr.FindOneByUUIDFunc = func(uuid uuid.UUID) (*entity.Computer, error) {
		return new(entity.Computer), nil
	}

	h := handler.NewUnassignComputerFromEmployee(er, cr)
	err := h.Unassign(cmd)

	assert.True(t, err != nil, "should fail with conflict error.")
}

func thenUnassignComputerFromEmployeeShouldUnassignEmployeeSuccessfully(
	t *testing.T,
	cr *mock.ComputerRepository,
	er *mock.EmployeeRepository,
	cmd *command.UnassignEmployeeFromComputer,
) {
	employee := new(entity.Employee)
	computer := &entity.Computer{
		ID:          1,
		UUID:        cmd.ComputerUUID,
		MACAddress:  "mac_address",
		IPAddress:   "ip_address",
		Name:        "name",
		Description: "description",
		Employee:    employee,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	cr.FindOneByUUIDFunc = func(uuid uuid.UUID) (*entity.Computer, error) {
		return computer, nil
	}

	cr.UpdateFunc = func(computer *entity.Computer) error {
		return nil
	}

	h := handler.NewUnassignComputerFromEmployee(er, cr)
	err := h.Unassign(cmd)
	if err != nil {
		t.FailNow()
	}

	assert.False(t, computer.HasEmployee(), "computer should not have an employee.")
}
