package mock

import (
	"github.com/google/uuid"
	"github.com/shonjord/tracker/internal/pkg/domain/entity"
)

type (
	EmployeeRepository struct {
		FindOneByUUIDFunc func(uuid.UUID) (*entity.Employee, error)
	}
)

// FindOneByUUID refer to the consumer of the interface for documentation.
func (m *EmployeeRepository) FindOneByUUID(uuid uuid.UUID) (*entity.Employee, error) {
	return m.FindOneByUUIDFunc(uuid)
}
