package handler_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/shonjord/tracker/internal/pkg/application/handler"
	"github.com/shonjord/tracker/internal/pkg/domain/entity"
	"github.com/shonjord/tracker/internal/pkg/domain/errors"
	"github.com/shonjord/tracker/internal/pkg/domain/query"
	assert "github.com/shonjord/tracker/test/.assert"
	mock "github.com/shonjord/tracker/test/.mock"
)

func TestGetComputerByUUID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		scenario string
		function func(
			*testing.T,
			*mock.ComputerRepository,
			*query.GetComputerByUUID,
		)
	}{
		{
			// Given a valid domain query
			scenario: "When computer repository can't find computer by its UUID",
			function: thenGetComputerByUUIDShouldReturnComputerNotFoundError,
		},
		{
			// Given a valid domain query
			scenario: "When computer repository returns the computer for the given UUID",
			function: thenGetComputerByUUIDShouldReturnNoError,
		},
	}

	qry := &query.GetComputerByUUID{
		ComputerUUID: uuid.New(),
	}

	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			test.function(
				t,
				new(mock.ComputerRepository),
				qry,
			)
		})
	}
}

func thenGetComputerByUUIDShouldReturnComputerNotFoundError(
	t *testing.T,
	cr *mock.ComputerRepository,
	qry *query.GetComputerByUUID,
) {
	cr.FindOneByUUIDFunc = func(uuid uuid.UUID) (*entity.Computer, error) {
		return nil, errors.NewComputerForUUIDNotFound(uuid)
	}

	h := handler.NewGetComputerByUUID(cr)
	_, err := h.GetComputer(qry)

	assert.True(t, err != nil, "should fail with computer repository errors.")
}

func thenGetComputerByUUIDShouldReturnNoError(
	t *testing.T,
	cr *mock.ComputerRepository,
	qry *query.GetComputerByUUID,
) {
	computer := new(entity.Computer)
	cr.FindOneByUUIDFunc = func(uuid uuid.UUID) (*entity.Computer, error) {
		return computer, nil
	}

	h := handler.NewGetComputerByUUID(cr)
	c, _ := h.GetComputer(qry)

	assert.Equals(t, computer, c)
}
