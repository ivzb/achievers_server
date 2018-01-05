package model

import (
	"time"
)

type QuestAchievement struct {
	ID string `json:"id"`

	QuestID       string `json:"quest_id"`
	AchievementID string `json:"achievement_id"`
	AuthorID      string `json:"author_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (db *DB) QuestAchievementExists(questID string, achievementID string) (bool, error) {
	return existsMultiple(db, "quest_achievement", []string{"quest_id", "achievement_id"}, []string{questID, achievementID})
}

func (db *DB) QuestAchievementSingle(qstID string, achID string) (*QuestAchievement, error) {
	qstAch := new(QuestAchievement)

	qstAch.QuestID = qstID
	qstAch.AchievementID = achID

	row := db.QueryRow("SELECT `id`, `author_id`, `created_at`, `updated_at`, `deleted_at` "+
		"FROM quest_achievement "+
		"WHERE quest_id = ? AND achievement_id = ? "+
		"LIMIT 1", qstID, achID)

	err := row.Scan(
		&qstAch.ID,
		&qstAch.AuthorID,
		&qstAch.CreatedAt,
		&qstAch.UpdatedAt,
		&qstAch.DeletedAt)

	if err != nil {
		return nil, err
	}

	return qstAch, nil
}

//func (db *DB) QuestsAll(page int) ([]*Quest, error) {
//offset := limit * page

//rows, err := db.Query("SELECT `id`, `title`, `picture_url`, `involvement_id`, `quest_type_id`, `author_id`, `created_at`, `updated_at`, `deleted_at` "+
//"FROM quest "+
//"ORDER BY `created_at` DESC "+
//"LIMIT ? OFFSET ?", limit, offset)

//if err != nil {
//return nil, err
//}

//defer rows.Close()

//qsts := make([]*Quest, 0)

//for rows.Next() {
//qst := new(Quest)
//err := rows.Scan(
//&qst.ID,
//&qst.Title,
//&qst.PictureURL,
//&qst.InvolvementID,
//&qst.QuestTypeID,
//&qst.AuthorID,
//&qst.CreatedAt,
//&qst.UpdatedAt,
//&qst.DeletedAt)

//if err != nil {
//return nil, err
//}

//qsts = append(qsts, qst)
//}

//if err = rows.Err(); err != nil {
//return nil, err
//}

//return qsts, nil
//}

//func (db *DB) QuestCreate(quest *Quest) (string, error) {
//return create(db, `INSERT INTO quest (id, title, picture_url, involvement_id, quest_type_id, author_id)
//VALUES(?, ?, ?, ?, ?, ?)`,
//quest.Title,
//quest.PictureURL,
//quest.InvolvementID,
//quest.QuestTypeID,
//quest.AuthorID)
//}
