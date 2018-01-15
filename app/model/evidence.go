package model

import "time"

type Evidence struct {
	ID string `json:"id"`

	Title      string `json:"title"`
	PictureURL string `json:"picture_url"`
	URL        string `json:"url"`

	MultimediaTypeID uint8  `json:"multimedia_type_id"`
	AchievementID    string `json:"achievement_id"`
	UserID           string `json:"user_id"`

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

	row := db.QueryRow("SELECT `title`, `picture_url`, `url`, `multimedia_type_id`, `achievement_id`, `user_id`, `created_at`, `updated_at`, `deleted_at` "+
		"FROM evidence "+
		"WHERE id = ? "+
		"LIMIT 1", id)

	err := row.Scan(
		&evd.Title,
		&evd.PictureURL,
		&evd.URL,
		&evd.MultimediaTypeID,
		&evd.AchievementID,
		&evd.UserID,
		&evd.CreatedAt,
		&evd.UpdatedAt,
		&evd.DeletedAt)

	if err != nil {
		return nil, err
	}

	return evd, nil
}

func (db *DB) EvidencesAll(page int) ([]*Evidence, error) {
	offset := limit * page

	rows, err := db.Query("SELECT `id`, `title`, `picture_url`, `url`, `multimedia_type_id`, `achievement_id`, `user_id`, `created_at`, `updated_at`, `deleted_at` "+
		"FROM evidence "+
		"ORDER BY `created_at` DESC "+
		"LIMIT ? OFFSET ?", limit, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	evds := make([]*Evidence, 0)

	for rows.Next() {
		evd := new(Evidence)
		err := rows.Scan(
			&evd.ID,
			&evd.Title,
			&evd.PictureURL,
			&evd.URL,
			&evd.MultimediaTypeID,
			&evd.AchievementID,
			&evd.UserID,
			&evd.CreatedAt,
			&evd.UpdatedAt,
			&evd.DeletedAt)

		if err != nil {
			return nil, err
		}

		evds = append(evds, evd)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return evds, nil
}

// EvidenceCreate saves evidence object to db
func (db *DB) EvidenceCreate(evidence *Evidence) (string, error) {
	return create(db, `INSERT INTO evidence (id, title, picture_url, url, multimedia_type_id, achievement_id, user_id)
        VALUES(?, ?, ?, ?, ?, ?, ?)`,
		evidence.Title,
		evidence.PictureURL,
		evidence.URL,
		evidence.MultimediaTypeID,
		evidence.AchievementID,
		evidence.UserID)
}
