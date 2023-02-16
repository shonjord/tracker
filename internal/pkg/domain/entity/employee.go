package entity

import (
	"time"

	"github.com/google/uuid"
)

type (
	Employee struct {
		ID           int         `json:"-"`
		UUID         uuid.UUID   `json:"uuid"`
		Name         string      `json:"name"`
		Abbreviation string      `json:"abbreviation"`
		Computers    []*Computer `json:"computers"`
		CreatedAt    time.Time   `json:"-"`
		UpdatedAt    time.Time   `json:"-"`
	}
)

// HasAbbreviation verifies if this employee has an abbreviation or not.
func (e *Employee) HasAbbreviation() bool {
	return "" != e.Abbreviation
}

func (e *Employee) HasComputers() bool {
	return nil != e.Computers || len(e.Computers) > 0
}

// HasBeenAssignedWithThreeOrMoreComputers verifies if this employee has been assigned
// with 3 or more computers.
func (e *Employee) HasBeenAssignedWithThreeOrMoreComputers() bool {
	return e.HasComputers() && len(e.Computers) >= 3
}

// AddComputers add multiple computers to this employee.
func (e *Employee) AddComputers(c []*Computer) {
	for _, computer := range c {
		e.AddComputer(computer)
	}
}

// AddComputer appends a new computer to the employee.
func (e *Employee) AddComputer(c *Computer) {
	e.Computers = append(e.Computers, c)
}
