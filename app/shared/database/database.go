package database

import (
	"fmt"
)

// Type is the type of database from a Type* constant
type Type string

const (
	// TypePostgre is Postgre
	TypePostgre Type = "Postgre"
)

// Info contains the database configurations
type Info struct {
	// Database type
	Type Type
	// Postgre info if used
	Postgre PostgreInfo
}

// PostgreInfo is the details for the database connection
type PostgreInfo struct {
	Username  string
	Password  string
	Name      string
	Hostname  string
	Port      int
	Parameter string
	PageLimit int
}

// DSN returns the Data Source Name
// postgres://username:password@localhost:port/db_name?sslmode=disable
func DSN(mi PostgreInfo) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?%s",
		mi.Username,
		mi.Password,
		mi.Hostname,
		mi.Port,
		mi.Name,
		mi.Parameter)
}
