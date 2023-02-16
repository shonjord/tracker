package handler

import (
	"github.com/shonjord/tracker/internal/pkg/domain/entity"
	"github.com/shonjord/tracker/internal/pkg/domain/query"
)

type (
	GetAllComputers struct {
		finder computerFinder
	}
)

// NewGetAllComputers returns a new instance of this handler.
func NewGetAllComputers(f computerFinder) *GetAllComputers {
	return &GetAllComputers{
		finder: f,
	}
}

// GetAllComputers returns all computers with a limit in case this is provided.
func (h *GetAllComputers) GetAllComputers(c *query.GetAllComputers) ([]*entity.Computer, error) {
	computers, err := h.finder.FindMany(c.Limit)
	if err != nil {
		return nil, err
	}

	return computers, nil
}
