package handler_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/shonjord/tracker/internal/pkg/application/handler"
	"github.com/shonjord/tracker/internal/pkg/domain/entity"
	"github.com/shonjord/tracker/internal/pkg/domain/errors"
	"github.com/shonjord/tracker/internal/pkg/domain/query"
	assert "github.com/shonjord/tracker/test/.assert"
	mock "github.com/shonjord/tracker/test/.mock"
)

func TestGetComputersAssignedToEmployee(t *testing.T) {
	t.Parallel()

	tests := []struct {
		scenario string
		function func(
			*testing.T,
			*mock.EmployeeRepository,
			*query.GetComputersAssignedToEmployee,
		)
	}{
		{
			// Given a valid domain query
			scenario: "When employee repository can't find employee by its UUID",
			function: thenGetComputersAssignedToEmployeeShouldReturnEmployeeNotFoundError,
		},
		{
			// Given a valid domain query
			scenario: "When employee repository finds the employee by its UUID",
			function: thenGetComputersAssignedToEmployeeShouldReturnComputersAssignedToEmployee,
		},
	}

	qry := &query.GetComputersAssignedToEmployee{
		EmployeeUUID: uuid.New(),
	}

	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			test.function(
				t,
				new(mock.EmployeeRepository),
				qry,
			)
		})
	}
}

func thenGetComputersAssignedToEmployeeShouldReturnEmployeeNotFoundError(
	t *testing.T,
	er *mock.EmployeeRepository,
	qry *query.GetComputersAssignedToEmployee,
) {
	er.FindOneByUUIDFunc = func(uuid uuid.UUID) (*entity.Employee, error) {
		return nil, errors.NewEmployeeForUUIDNotFound(uuid)
	}

	h := handler.NewGetComputersAssignedToEmployee(er)
	_, err := h.GetComputersAssignedToEmployee(qry)

	assert.True(t, err != nil, "should fail with employee repository errors.")
}

func thenGetComputersAssignedToEmployeeShouldReturnComputersAssignedToEmployee(
	t *testing.T,
	er *mock.EmployeeRepository,
	qry *query.GetComputersAssignedToEmployee,
) {
	employee := &entity.Employee{
		ID:           1,
		UUID:         qry.EmployeeUUID,
		Name:         "employee",
		Abbreviation: "abv",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	computer := &entity.Computer{
		ID:          1,
		UUID:        uuid.New(),
		MACAddress:  "mac_address",
		IPAddress:   "ip_address",
		Name:        "name",
		Description: "description",
		Employee:    employee,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	computers := []*entity.Computer{computer}

	employee.AddComputers(computers)

	er.FindOneByUUIDFunc = func(uuid uuid.UUID) (*entity.Employee, error) {
		return employee, nil
	}

	h := handler.NewGetComputersAssignedToEmployee(er)
	c, _ := h.GetComputersAssignedToEmployee(qry)

	assert.Equals(t, computers, c)
}
