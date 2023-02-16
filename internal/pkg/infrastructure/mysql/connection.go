package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

type (
	Connection struct {
		connection *sql.DB
	}
)

// NewConnection returns a new DB connection after checking that is reachable.
func NewConnection(conn *sql.DB) *Connection {
	return &Connection{
		connection: conn,
	}
}

// FindOneBy scans one struct.
func (d *Connection) FindOneBy(query string, dest ...interface{}) error {
	return d.connection.QueryRow(query).Scan(dest...)
}

// FindManyBy scans a callback and executes the given query.
func (d *Connection) FindManyBy(query string, callback func(*sql.Rows) error) error {
	rows, err := d.connection.Query(query)
	if err != nil {
		return err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.WithError(err).Error("errors while closing MySQL rows")
		}
	}()

	for rows.Next() {
		if err = callback(rows); err != nil {
			return err
		}
	}

	if err = rows.Err(); err != nil {
		return err
	}

	return nil
}

// Prepare prepares a statement for the given query.
func (d *Connection) Prepare(q string) (*sql.Stmt, error) {
	return d.connection.Prepare(q)
}

// Execute executes a void statement.
func (d *Connection) Execute(query string, args ...interface{}) error {
	_, err := d.connection.Exec(query, args...)

	return err
}
