package handler_test

import (
	"errors"
	"testing"

	"github.com/shonjord/tracker/internal/pkg/application/handler"
	"github.com/shonjord/tracker/internal/pkg/domain/entity"
	"github.com/shonjord/tracker/internal/pkg/domain/query"
	assert "github.com/shonjord/tracker/test/.assert"
	mock "github.com/shonjord/tracker/test/.mock"
)

func TestGetAllComputers(t *testing.T) {
	t.Parallel()

	tests := []struct {
		scenario string
		function func(
			*testing.T,
			*mock.ComputerRepository,
			*query.GetAllComputers,
		)
	}{
		{
			// Given a valid domain query
			scenario: "When computer repository can't find all computers",
			function: thenGetAllComputersShouldReturnInternalError,
		},
		{
			// Given a valid domain query
			scenario: "When computer repository finds all computers",
			function: thenGetAllComputersShouldReturnAllComputers,
		},
	}

	qry := &query.GetAllComputers{
		Limit: 0,
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

func thenGetAllComputersShouldReturnInternalError(
	t *testing.T,
	cr *mock.ComputerRepository,
	qry *query.GetAllComputers,
) {
	cr.FindManyFunc = func(limit int) ([]*entity.Computer, error) {
		return nil, errors.New("SQL error")
	}

	h := handler.NewGetAllComputers(cr)
	_, err := h.GetAllComputers(qry)

	assert.True(t, err != nil, "should fail with computer repository errors.")
}

func thenGetAllComputersShouldReturnAllComputers(
	t *testing.T,
	cr *mock.ComputerRepository,
	qry *query.GetAllComputers,
) {
	computer := new(entity.Computer)

	cr.FindManyFunc = func(limit int) ([]*entity.Computer, error) {
		return []*entity.Computer{computer}, nil
	}

	h := handler.NewGetAllComputers(cr)
	computers, _ := h.GetAllComputers(qry)

	assert.Equals(t, computer, computers[0])
}
