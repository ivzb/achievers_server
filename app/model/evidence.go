package model

import "time"

type Evidence struct {
	ID string `json:"id"`

	Description string `json:"description"`
	PreviewURL  string `json:"preview_url"`
	URL         string `json:"url"`

	MultimediaTypeID uint8  `json:"multimedia_type_id"`
	AchievementID    string `json:"achievement_id"`
	AuthorID         string `json:"author_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (db *DB) EvidenceExists(id string) (bool, error) {
	return exists(db, "evidence", "id", id)
}

func (db *DB) EvidenceSingle(id string) (*Evidence, error) {
	evd := new(Evidence)

	evd.ID = id

	row := db.QueryRow("SELECT `description`, `preview_url`, `url`, `multimedia_type_id`, `achievement_id`, `author_id`, `created_at`, `updated_at`, `deleted_at` "+
		"FROM evidence "+
		"WHERE id = ? "+
		"LIMIT 1", id)

	err := row.Scan(
		&evd.Description,
		&evd.PreviewURL,
		&evd.URL,
		&evd.MultimediaTypeID,
		&evd.AchievementID,
		&evd.AuthorID,
		&evd.CreatedAt,
		&evd.UpdatedAt,
		&evd.DeletedAt)

	if err != nil {
		return nil, err
	}

	return evd, nil
}

// EvidenceCreate saves evidence object to db
func (db *DB) EvidenceCreate(evidence *Evidence) (string, error) {
	return create(db, `INSERT INTO evidence (id, description, preview_url, url, multimedia_type_id, achievement_id, author_id)
        VALUES(?, ?, ?, ?, ?, ?, ?)`,
		evidence.Description,
		evidence.PreviewURL,
		evidence.URL,
		evidence.MultimediaTypeID,
		evidence.AchievementID,
		evidence.AuthorID)
}
