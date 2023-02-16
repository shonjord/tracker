package handler

import (
	"github.com/shonjord/tracker/internal/pkg/domain/entity"
	"github.com/shonjord/tracker/internal/pkg/domain/query"
)

type (
	GetComputerByUUID struct {
		finder computerFinder
	}
)

// NewGetComputerByUUID returns a new instance of this handler.
func NewGetComputerByUUID(f computerFinder) *GetComputerByUUID {
	return &GetComputerByUUID{
		finder: f,
	}
}

// GetComputer returns a computer for the given identifier.
func (h *GetComputerByUUID) GetComputer(q *query.GetComputerByUUID) (*entity.Computer, error) {
	computer, err := h.finder.FindOneByUUID(q.ComputerUUID)
	if err != nil {
		return nil, err
	}

	return computer, nil
}
