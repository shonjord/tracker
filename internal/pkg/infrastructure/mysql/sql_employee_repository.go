package mysql

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/shonjord/tracker/internal/pkg/domain/entity"
	"github.com/shonjord/tracker/internal/pkg/domain/errors"
)

type (
	EmployeeRepository struct {
		connection         *Connection
		computerRepository *ComputerRepository
	}
)

// NewEmployeeRepository returns a new instance of this repository.
func NewEmployeeRepository(c *Connection) *EmployeeRepository {
	return &EmployeeRepository{
		connection: c,
	}
}

// WithComputerRepository assigns computer repository to this repository.
func (r *EmployeeRepository) WithComputerRepository(cr *ComputerRepository) {
	r.computerRepository = cr
}

// FindOneByUUID finds one employee by the given UUID.
func (r *EmployeeRepository) FindOneByUUID(uuid uuid.UUID) (*entity.Employee, error) {
	var (
		employee = new(entity.Employee)
		query    = fmt.Sprintf("SELECT * FROM employees WHERE uuid = '%s'", uuid)
	)

	if err := r.connection.FindOneBy(query, employeeDbValues(employee)...); err != nil {
		if errorIsNoRows(err) {
			return nil, errors.NewEmployeeForUUIDNotFound(uuid)
		}
		return nil, err
	}

	computers, err := r.computerRepository.FindManyForEmployeeWithUUID(employee.UUID)
	if err != nil {
		return nil, err
	}

	if len(computers) > 0 {
		employee.AddComputers(computers)
	}

	return employee, nil
}
