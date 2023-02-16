package entity

import (
	"time"

	"github.com/google/uuid"
)

type (
	Computer struct {
		ID          int       `json:"-"`
		UUID        uuid.UUID `json:"uuid"`
		MACAddress  string    `json:"MACAddress"`
		IPAddress   string    `json:"IPAddress"`
		Name        string    `json:"Name"`
		Description string    `json:"Description"`
		Employee    *Employee `json:"-"`
		CreatedAt   time.Time `json:"-"`
		UpdatedAt   time.Time `json:"-"`
	}
)

// NewComputer returns a new instance of a computer.
func NewComputer(ma string, ipa string, n string, d string, e *Employee) *Computer {
	return &Computer{
		UUID:        uuid.New(),
		MACAddress:  ma,
		IPAddress:   ipa,
		Name:        n,
		Description: d,
		Employee:    e,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

// WithEmployee assigns an employee to this computer.
func (c *Computer) WithEmployee(e *Employee) {
	c.Employee = e
	e.AddComputer(c)
}

// HasEmployee verifies if this computer has been assigned to an employee or not.
func (c *Computer) HasEmployee() bool {
	return nil != c.Employee
}

// UnassignEmployee removes the current employee from this computer's state.
func (c *Computer) UnassignEmployee() {
	c.Employee = nil
}
