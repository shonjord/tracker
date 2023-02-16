package mysql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shonjord/tracker/internal/pkg/domain/entity"
	"github.com/shonjord/tracker/internal/pkg/domain/errors"
)

const (
	emptyUUID = "00000000-0000-0000-0000-000000000000"
)

type (
	ComputerRepository struct {
		connection         *Connection
		employeeRepository *EmployeeRepository
	}
)

// NewComputerRepository returns a new instance of this repository.
func NewComputerRepository(c *Connection) *ComputerRepository {
	return &ComputerRepository{
		connection: c,
	}
}

func (r *ComputerRepository) WithEmployeeRepository(er *EmployeeRepository) {
	r.employeeRepository = er
}

// Save persist a new trip into MySQL DB.
func (r *ComputerRepository) Save(e *entity.Computer) error {
	query := "INSERT INTO computers (uuid, employee_uuid, mac_address, ip_address, name, description, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	smt, err := r.connection.Prepare(query)
	if err != nil {
		return err
	}

	values := []interface{}{
		e.UUID,
		e.Employee.UUID,
		e.MACAddress,
		e.IPAddress,
		e.Name,
		e.Description,
		e.CreatedAt,
		e.UpdatedAt,
	}

	_, err = smt.Exec(values...)
	if err != nil {
		if errorIsConflict(err) {
			return errors.NewMacAddressIsAssignedToAnotherEmployee(e)
		}
		return err
	}

	return nil
}

// FindOneByUUID finds one computer by the given UUID.
func (r *ComputerRepository) FindOneByUUID(uuid uuid.UUID) (*entity.Computer, error) {
	var (
		computer = new(entity.Computer)
		employee = new(entity.Employee)
		query    = fmt.Sprintf("SELECT * FROM computers WHERE uuid = '%s'", uuid)
	)

	if err := r.connection.FindOneBy(query, computerDbValues(computer, employee)...); err != nil {
		if errorIsNoRows(err) {
			return nil, errors.NewComputerForUUIDNotFound(uuid)
		}
		return nil, err
	}

	if emptyUUID == employee.UUID.String() {
		return computer, nil
	}

	employee, err := r.employeeRepository.FindOneByUUID(employee.UUID)
	if err != nil {
		return nil, err
	}

	computer.WithEmployee(employee)

	return computer, nil
}

// FindManyForEmployeeWithUUID returns a collection of computers assigned to a specific employee.
func (r *ComputerRepository) FindManyForEmployeeWithUUID(uuid uuid.UUID) ([]*entity.Computer, error) {
	var (
		employee  = new(entity.Employee)
		computers []*entity.Computer
		query     = fmt.Sprintf("SELECT * FROM computers WHERE employee_uuid = '%s'", uuid)
	)

	callback := func(rows *sql.Rows) error {
		computer := new(entity.Computer)

		if err := rows.Scan(computerDbValues(computer, employee)...); err != nil {
			return err
		}

		computers = append(computers, computer)

		return nil
	}

	if err := r.connection.FindManyBy(query, callback); err != nil {
		return nil, err
	}

	return computers, nil
}

// FindMany returns a collection of computers with a limit in case this is provided.
func (r *ComputerRepository) FindMany(limit int) ([]*entity.Computer, error) {
	var (
		employee  = new(entity.Employee)
		computers []*entity.Computer
		query     = fmt.Sprint("SELECT * FROM computers")
	)

	if 0 != limit {
		query = fmt.Sprintf("%s LIMIT %d", query, limit)
	}

	callback := func(rows *sql.Rows) error {
		computer := new(entity.Computer)

		if err := rows.Scan(computerDbValues(computer, employee)...); err != nil {
			return err
		}

		computers = append(computers, computer)

		return nil
	}

	if err := r.connection.FindManyBy(query, callback); err != nil {
		return nil, err
	}

	return computers, nil
}

// Update updates all values for the given computer.
func (r *ComputerRepository) Update(e *entity.Computer) error {
	var (
		query = "UPDATE computers SET mac_address = ?, ip_address = ?, name = ?, description = ?, employee_uuid = ?, updated_at = ? WHERE id = ?"
	)

	var employeeUUID *uuid.UUID

	if !e.HasEmployee() {
		employeeUUID = nil
	} else {
		employeeUUID = &e.Employee.UUID
	}

	return r.connection.Execute(
		query,
		e.MACAddress,
		e.IPAddress,
		e.Name,
		e.Description,
		employeeUUID,
		time.Now(),
		e.ID,
	)
}
