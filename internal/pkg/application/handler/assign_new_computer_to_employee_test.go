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

func TestAssignNewComputerToEmployee(t *testing.T) {
	t.Parallel()

	tests := []struct {
		scenario string
		function func(
			*testing.T,
			*mock.ComputerRepository,
			*mock.EmployeeRepository,
			*mock.AdminNotificationClient,
			*command.AssignNewComputerToEmployee,
		)
	}{
		{
			// Given a valid domain command
			scenario: "When employee repository can't find an employee",
			function: thenAssignNewComputerToEmployeeShouldReturnEmployeeNotFoundError,
		},
		{
			// Given a valid domain command
			scenario: "When computer repository can't save a new computer",
			function: thenAssignNewComputerToEmployeeShouldReturnInternalError,
		},
		{
			// Given a valid domain command
			scenario: "When admin notification client notifies",
			function: thenAssignNewComputerToEmployeeShouldReturnNoError,
		},
	}

	cmd := &command.AssignNewComputerToEmployee{
		MacAddress:   "mac_address",
		IPAddress:    "ip_address",
		Name:         "name",
		Description:  "description",
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

func thenAssignNewComputerToEmployeeShouldReturnEmployeeNotFoundError(
	t *testing.T,
	cr *mock.ComputerRepository,
	er *mock.EmployeeRepository,
	nc *mock.AdminNotificationClient,
	cmd *command.AssignNewComputerToEmployee,
) {
	er.FindOneByUUIDFunc = func(u uuid.UUID) (*entity.Employee, error) {
		return nil, errors.New("employee not found")
	}

	h := handler.NewAssignNewComputerToEmployee(er, cr, nc)
	_, err := h.Assign(cmd)

	assert.True(t, err != nil, "should fail with employee repository errors.")
}

func thenAssignNewComputerToEmployeeShouldReturnInternalError(
	t *testing.T,
	cr *mock.ComputerRepository,
	er *mock.EmployeeRepository,
	nc *mock.AdminNotificationClient,
	cmd *command.AssignNewComputerToEmployee,
) {
	er.FindOneByUUIDFunc = func(u uuid.UUID) (*entity.Employee, error) {
		return new(entity.Employee), nil
	}

	cr.SaveFunc = func(computer *entity.Computer) error {
		return errors.New("SQL error")
	}

	h := handler.NewAssignNewComputerToEmployee(er, cr, nc)
	_, err := h.Assign(cmd)

	assert.True(t, err != nil, "should fail with computer repository errors.")
}

func thenAssignNewComputerToEmployeeShouldReturnNoError(
	t *testing.T,
	cr *mock.ComputerRepository,
	er *mock.EmployeeRepository,
	nc *mock.AdminNotificationClient,
	cmd *command.AssignNewComputerToEmployee,
) {
	er.FindOneByUUIDFunc = func(u uuid.UUID) (*entity.Employee, error) {
		return new(entity.Employee), nil
	}

	cr.SaveFunc = func(computer *entity.Computer) error {
		return nil
	}

	nc.NotifyFunc = func(computer *entity.Computer) error {
		assert.True(t, true, "notification service should notify at this point")

		return nil
	}

	h := handler.NewAssignNewComputerToEmployee(er, cr, nc)
	_, err := h.Assign(cmd)

	assert.True(t, err == nil, "handler should finish successfully.")
}
