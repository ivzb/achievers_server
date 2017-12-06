package auth

import (
	"fmt"
    
    "app/model/user"
	"app/shared/database"
)

const tableName = "user"

type Entity struct {
	Email     string    `db:"email" json:"email" require:"true"`
	Password  string    `db:"password" json:"password" require:"true"`
}

func New() (*Entity) {
	entity := &Entity{}

	return entity
}

func (e *Entity) Auth() (*user.Entity, error) {
	result := &user.Entity{}

	err := database.SQL.Get(
		result, 
		fmt.Sprintf("SELECT id FROM %v WHERE email = ? AND password = ? LIMIT 1", tableName), 
		e.Email, 
		e.Password)

	return result, err
}