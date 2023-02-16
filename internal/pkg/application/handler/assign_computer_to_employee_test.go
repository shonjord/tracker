package handler_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/shonjord/tracker/internal/pkg/application/handler"
	"github.com/shonjord/tracker/internal/pkg/domain/command"
	"github.com/shonjord/tracker/internal/pkg/domain/entity"
	assert "github.com/shonjord/tracker/test/.assert"
	mock "github.com/shonjord/tracker/test/.mock"
)

func TestAssignComputerToEmployee(t *testing.T) {
	t.Parallel()

	tests := []struct {
		scenario string
		function func(
			*testing.T,
			*mock.ComputerRepository,
			*mock.EmployeeRepository,
			*mock.AdminNotificationClient,
			*command.AssignComputerToEmployee,
		)
	}{
		{
			// Given a valid domain command
			scenario: "When computer repository can't find a computer",
			function: thenAssignComputerToEmployeeShouldReturnComputerNotFoundError,
		},
		{
			// Given a valid domain command
			scenario: "When employee repository can't find an employee",
			function: thenAssignComputerToEmployeeShouldReturnEmployeeNotFoundError,
		},
		{
			// Given a valid domain command
			scenario: "When computer repository can't find update",
			function: thenAssignComputerToEmployeeShouldReturnInternalError,
		},
		{
			// Given a valid domain command
			scenario: "When admin notification client notifies",
			function: thenAssignComputerToEmployeeShouldReturnNoError,
		},
	}

	cmd := &command.AssignComputerToEmployee{
		ComputerUUID: uuid.New(),
		EmployeeUUID: uuid.New(),
	}

	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			test.function(
				t,
				new(mock.ComputerRepository),
				new(mock.EmployeeRepository),
				new(mock.AdminNotificationClient),
				cmd,
			)
		})
	}
}

func thenAssignComputerToEmployeeShouldReturnComputerNotFoundError(
	t *testing.T,
	cr *mock.ComputerRepository,
	er *mock.EmployeeRepository,
	nc *mock.AdminNotificationClient,
	cmd *command.AssignComputerToEmployee,
) {
	cr.FindOneByUUIDFunc = func(u uuid.UUID) (*entity.Computer, error) {
		return nil, errors.New("computer not found")
	}

	h := handler.NewAssignComputerToEmployee(er, cr, nc)
	err := h.Assign(cmd)

	assert.True(t, err != nil, "should fail with computer repository errors.")
}

func thenAssignComputerToEmployeeShouldReturnEmployeeNotFoundError(
	t *testing.T,
	cr *mock.ComputerRepository,
	er *mock.EmployeeRepository,
	nc *mock.AdminNotificationClient,
	cmd *command.AssignComputerToEmployee,
) {
	cr.FindOneByUUIDFunc = func(u uuid.UUID) (*entity.Computer, error) {
		return new(entity.Computer), nil
	}

	er.FindOneByUUIDFunc = func(u uuid.UUID) (*entity.Employee, error) {
		return nil, errors.New("employee not found")
	}

	h := handler.NewAssignComputerToEmployee(er, cr, nc)
	err := h.Assign(cmd)

	assert.True(t, err != nil, "should fail with employee repository errors.")
}

func thenAssignComputerToEmployeeShouldReturnInternalError(
	t *testing.T,
	cr *mock.ComputerRepository,
	er *mock.EmployeeRepository,
	nc *mock.AdminNotificationClient,
	cmd *command.AssignComputerToEmployee,
) {
	cr.FindOneByUUIDFunc = func(u uuid.UUID) (*entity.Computer, error) {
		return new(entity.Computer), nil
	}

	er.FindOneByUUIDFunc = func(u uuid.UUID) (*entity.Employee, error) {
		return new(entity.Employee), nil
	}

	cr.UpdateFunc = func(computer *entity.Computer) error {
		return errors.New("SQL error")
	}

	h := handler.NewAssignComputerToEmployee(er, cr, nc)
	err := h.Assign(cmd)

	assert.True(t, err != nil, "should fail with computer repository errors.")
}

func thenAssignComputerToEmployeeShouldReturnNoError(
	t *testing.T,
	cr *mock.ComputerRepository,
	er *mock.EmployeeRepository,
	nc *mock.AdminNotificationClient,
	cmd *command.AssignComputerToEmployee,
) {
	cr.FindOneByUUIDFunc = func(u uuid.UUID) (*entity.Computer, error) {
		return new(entity.Computer), nil
	}

	er.FindOneByUUIDFunc = func(u uuid.UUID) (*entity.Employee, error) {
		return new(entity.Employee), nil
	}

	cr.UpdateFunc = func(computer *entity.Computer) error {
		return nil
	}

	nc.NotifyFunc = func(computer *entity.Computer) error {
		assert.True(t, true, "notification service should notify at this point")

		return nil
	}

	h := handler.NewAssignComputerToEmployee(er, cr, nc)
	err := h.Assign(cmd)

	assert.True(t, err == nil, "handler should finish successfully.")
}
