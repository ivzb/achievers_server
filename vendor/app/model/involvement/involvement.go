package involvement

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"app/shared/database"
)

// Errors
var (
	ErrNoResult = errors.New("no result")
	ErrExists   = errors.New("already exists")
	ErrNotExist = errors.New("does not exist")
)

// Name of the table
const tableName = "involvement"

// Entity information
type Entity struct {
	ID        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name" require:"true"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt time.Time `db:"deleted_at" json:"deleted_at"`
}

// Group of entities
type Group []Entity

// *****************************************************************************
// Read
// *****************************************************************************

// Read returns one entity with the matching ID
// If no result, it will return sql.ErrNoRows
func Read(ID string) (*Entity, error) {
	return readOneByField("id", ID)
}

// ReadAll returns all entities
func ReadAll() (Group, error) {
	var result Group
	err := database.SQL.Select(&result, fmt.Sprintf("SELECT * FROM %v", tableName))
	return result, err
}

// readOneByField returns the entity that matches the field value
// If no result, it will return ErrNoResult
func readOneByField(name string, value string) (*Entity, error) {
	result := &Entity{}
	err := database.SQL.Get(result, fmt.Sprintf("SELECT * FROM %v WHERE %v = ? LIMIT 1", tableName, name), value)
	if err == sql.ErrNoRows {
		err = ErrNoResult
	}
	return result, err
}

// readAllByField returns entities matching a field value
// If no result, it will return an empty group
func readAllByField(name string, value string) (Group, error) {
	var result Group
	err := database.SQL.Select(&result, fmt.Sprintf("SELECT * FROM %v WHERE %v = ?", tableName, name), value)
	return result, err
}

// *****************************************************************************
// Exist
// *****************************************************************************

// Read returns one entity with the matching ID
// If no result, it will return sql.ErrNoRows
func Exist(id string) (bool, error) {
	// return existOneByField("id", ID)
	u, err := readOneByField("id", id)

	if err != nil {
		return false, err
	}

	return u != nil, nil
}
