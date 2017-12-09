package achievement

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
	ErrNotExist = errors.New("does not exist")
)

// Name of the table
const tableName = "achievement"

// Entity information
type Entity struct {
	ID            string    `db:"id" json:"id"`
	Title         string    `db:"title" json:"title" require:"true"`
	Description   string    `db:"description" json:"description" require:"true"`
	PictureUrl    string    `db:"picture_url" json:"picture_url" require:"true"`
	InvolvementID string    `db:"involvement_id" json:"involvement_id" require:"true"`
	AuthorId      string    `db:"author_id" json:"author_id" require:"true"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt     time.Time `db:"deleted_at" json:"deleted_at"`
}

// Group of entities
type Group []Entity

// New entity
func New() (*Entity, error) {
	var err error
	entity := &Entity{}

	// Set the default parameters
	entity.ID, err = database.UUID()
	// If error on UUID generation
	if err != nil {
		return entity, err
	}

	return entity, nil
}

// *****************************************************************************
// Create
// *****************************************************************************

// Create will add a new entity
func (a *Entity) Create(uID string) (int, error) {
	// Create the entity
	_, err := database.SQL.Exec(fmt.Sprintf(`
		INSERT INTO %v
		(id, title, description, picture_url, involvement_id, author_id)
		VALUES
		(?,?,?,?,?, ?)
		`, tableName),
		a.ID,
		a.Title,
		a.Description,
		a.PictureUrl,
		a.InvolvementID,
		uID)

	// If error occurred error
	if err != nil {
		return 0, err
	}

	return 1, nil
}

// *****************************************************************************
// Read
// *****************************************************************************

// Read returns one entity with the matching ID
// If no result, it will return sql.ErrNoRows
func Get(ID string) (*Entity, error) {
	return readOneByField("id", ID)
}

// ReadAll returns all entities
func Load(page int) (Group, error) {
	var result Group

	pageSize := 9
	start := pageSize * page
	end := start + pageSize

	err := database.SQL.Select(&result, fmt.Sprintf("SELECT * FROM %v ORDER BY `updated_at` DESC LIMIT %v, %v", tableName, start, end))
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
