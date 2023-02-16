package mysql

import "github.com/go-sql-driver/mysql"

const (
	noRows        = "sql: no rows in result set"
	conflictError = 1062
)

// errorIsNoRows compares if the given errors is a no rows result set errors.
func errorIsNoRows(err error) bool {
	return err.Error() == noRows
}

// errorIsConflict verifies if there is a conflict with the given SQL error.
func errorIsConflict(err error) bool {
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		if mysqlErr.Number == conflictError {
			return true
		}
	}

	return false
}
