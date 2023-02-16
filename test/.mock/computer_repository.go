package mock

import (
	"github.com/google/uuid"
	"github.com/shonjord/tracker/internal/pkg/domain/entity"
)

type (
	ComputerRepository struct {
		SaveFunc                        func(*entity.Computer) error
		FindOneByUUIDFunc               func(uuid.UUID) (*entity.Computer, error)
		FindManyForEmployeeWithUUIDFunc func(uuid.UUID) ([]*entity.Computer, error)
		FindManyFunc                    func(int) ([]*entity.Computer, error)
		UpdateFunc                      func(*entity.Computer) error
	}
)

// Save refer to the consumer of the interface for documentation.
func (m *ComputerRepository) Save(c *entity.Computer) error {
	return m.SaveFunc(c)
}

// FindOneByUUID refer to the consumer of the interface for documentation.
func (m *ComputerRepository) FindOneByUUID(uuid uuid.UUID) (*entity.Computer, error) {
	return m.FindOneByUUIDFunc(uuid)
}

// FindManyForEmployeeWithUUID refer to the consumer of the interface for documentation.
func (m *ComputerRepository) FindManyForEmployeeWithUUID(uuid uuid.UUID) ([]*entity.Computer, error) {
	return m.FindManyForEmployeeWithUUIDFunc(uuid)
}

// FindMany refer to the consumer of the interface for documentation.
func (m *ComputerRepository) FindMany(limit int) ([]*entity.Computer, error) {
	return m.FindManyFunc(limit)
}

// Update refer to the consumer of the interface for documentation.
func (m *ComputerRepository) Update(e *entity.Computer) error {
	return m.UpdateFunc(e)
}
