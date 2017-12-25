package database

import (
	"fmt"
)

// Type is the type of database from a Type* constant
type Type string

const (
	// TypeMySQL is MySQL
	TypeMySQL Type = "MySQL"
)

// Info contains the database configurations
type Info struct {
	// Database type
	Type Type
	// MySQL info if used
	MySQL MySQLInfo
}

// MySQLInfo is the details for the database connection
type MySQLInfo struct {
	Username  string
	Password  string
	Name      string
	Hostname  string
	Port      int
	Parameter string
}

// DSN returns the Data Source Name
// Example: root:123@tcp(localhost:3306)/test?parseTime=true
func DSN(mi MySQLInfo) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s%s",
		mi.Username,
		mi.Password,
		mi.Hostname,
		mi.Port,
		mi.Name,
		mi.Parameter)
}

// AffectedRows returns the number of rows affected by the query
// Will panic if result does not exist
// func AffectedRows(result sql.Result) int {
// 	// If successful, get the number of affected rows
// 	count, err := result.RowsAffected()
// 	if err != nil { // Feature not supported
// 		// Only show error for admin
// 		log.Println(err)
// 		return 1
// 	}

// 	return int(count)
// }