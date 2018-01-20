package model

import (
	"time"
)

type Profile struct {
	ID string `json:"id"`

	Name string `json:"name"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (db *db.Profile) ProfileExists(id string) (bool, error) {
	return exists(db, "profile", "id", id)
}

func (db *db.Profile) ProfileSingle(id string) (*Profile, error) {
	prfl := new(Profile)

	prfl.ID = id

	row := db.QueryRow("SELECT `name`, `created_at`, `updated_at`, `deleted_at` "+
		"FROM profile "+
		"WHERE id = ? "+
		"LIMIT 1", id)

	err := row.Scan(
		&prfl.Name,
		&prfl.CreatedAt,
		&prfl.UpdatedAt,
		&prfl.DeletedAt)

	if err != nil {
		return nil, err
	}

	return prfl, nil
}

func (db *db.Profile) ProfileByUserID(userID string) (*Profile, error) {
	prfl := new(Profile)

	row := db.QueryRow("SELECT `id`, `name`, `created_at`, `updated_at`, `deleted_at` "+
		"FROM profile "+
		"WHERE user_id = ? "+
		"LIMIT 1", userID)

	err := row.Scan(
		&prfl.ID,
		&prfl.Name,
		&prfl.CreatedAt,
		&prfl.UpdatedAt,
		&prfl.DeletedAt)

	if err != nil {
		return nil, err
	}

	return prfl, nil
}

func (db *db.Profile) ProfileCreate(profile *Profile, userID string) (string, error) {
	return create(db, `INSERT INTO profile (id, name, user_id)
        VALUES(?, ?, ?)`,
		profile.Name,
		userID)
}
