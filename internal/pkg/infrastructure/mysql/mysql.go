package mysql

import (
	"database/sql"

	"github.com/shonjord/tracker/internal/pkg/domain/entity"
)

// computerDbValues returns all DB values of the computer entity.
func computerDbValues(c *entity.Computer, e *entity.Employee) []interface{} {
	return []interface{}{
		&c.ID,
		&c.UUID,
		&e.UUID,
		&c.MACAddress,
		&c.IPAddress,
		&c.Name,
		&c.Description,
		&c.CreatedAt,
		&c.UpdatedAt,
	}
}

// employeeDbValues returns all DB values of the employee entity.
func employeeDbValues(e *entity.Employee) []interface{} {
	var abbreviation sql.NullString

	return []interface{}{
		&e.ID,
		&e.UUID,
		&e.Name,
		&abbreviation,
		&e.CreatedAt,
		&e.UpdatedAt,
	}
}
