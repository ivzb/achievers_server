package model

import (
	"crypto/rand"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/ivzb/achievers_server/app/shared/database"

	// MySQL DB driver
	_ "github.com/go-sql-driver/mysql"
)

const (
	limit = 9
)

// DBSourcer contains all available DAO functions
type DBSourcer interface {
	UUID() (string, error)

	UserExists(id string) (bool, error)
	UserEmailExists(email string) (bool, error)
	UserCreate(user *User) (string, error)
	UserAuth(auth *Auth) (string, error)

	ProfileExists(id string) (bool, error)
	ProfileSingle(id string) (*Profile, error)
	ProfileByUserID(userID string) (*Profile, error)
	ProfileCreate(profile *Profile, userID string) (string, error)

	AchievementExists(id string) (bool, error)
	AchievementSingle(id string) (*Achievement, error)
	AchievementsLastID() (string, error)
	AchievementsAfter(afterID string) ([]*Achievement, error)
	AchievementsByQuestIDAfter(questID string, afterID string) ([]*Achievement, error)
	AchievementsByQuestIDLastID(questID string) (string, error)
	AchievementCreate(achievement *Achievement) (string, error)

	EvidenceExists(id string) (bool, error)
	EvidenceSingle(id string) (*Evidence, error)
	EvidencesAll(page int) ([]*Evidence, error)
	EvidenceCreate(evidence *Evidence) (string, error)

	QuestExists(id string) (bool, error)
	QuestSingle(id string) (*Quest, error)
	QuestsAll(page int) ([]*Quest, error)
	QuestCreate(quest *Quest) (string, error)

	QuestTypeExists(id uint8) (bool, error)

	QuestAchievementExists(questID string, achievementID string) (bool, error)
	QuestAchievementCreate(qstAch *QuestAchievement) (string, error)

	RewardExists(id string) (bool, error)
	RewardSingle(id string) (*Reward, error)
	RewardsAll(page int) ([]*Reward, error)
	RewardCreate(reward *Reward) (string, error)

	RewardTypeExists(id uint8) (bool, error)

	InvolvementExists(id string) (bool, error)

	MultimediaTypeExists(id uint8) (bool, error)
}

// DB struct holds the connection to DB
type DB struct {
	*sql.DB
}

// NewDB creates connection to the database
func NewDB(d database.Info) (*DB, error) {
	switch d.Type {
	case database.TypeMySQL:
		db, err := sql.Open("mysql", database.DSN(d.MySQL))
		if err != nil {
			return nil, err
		}
		if err = db.Ping(); err != nil {
			return nil, err
		}
		return &DB{db}, nil
	default:
		return nil, errors.New("No registered database in config")
	}
}

// UUID generates UUID for use as an ID
func (db *DB) UUID() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:]), nil
}

// exists checks whether row in specified table exists by column and value
func exists(db *DB, table string, column string, value string) (bool, error) {
	query := fmt.Sprintf("SELECT COUNT(id) FROM %s WHERE %s = ? LIMIT 1", table, column)
	stmt, err := db.Prepare(query)

	if err != nil {
		return false, err
	}

	var count int
	err = stmt.QueryRow(value).Scan(&count)

	if err != nil {
		return false, err
	}

	return count != 0, nil
}

// existsMultiple checks whether row in specified table exists by []columns and []values
func existsMultiple(db *DB, table string, columns []string, values []string) (bool, error) {
	query := fmt.Sprintf("SELECT COUNT(id) FROM %s WHERE %s LIMIT 1", table, whereClause(columns))
	stmt, err := db.Prepare(query)

	if err != nil {
		return false, err
	}

	var count int
	err = stmt.QueryRow(scanArgs(values)...).Scan(&count)

	if err != nil {
		return false, err
	}

	return count != 0, nil
}

func scanArgs(values []string) []interface{} {
	scanArgs := make([]interface{}, len(values))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	return scanArgs
}

func whereClause(columns []string) string {
	placeholders := make([]string, 0, len(columns))

	for _, column := range columns {
		placeholders = append(placeholders, column+" = ?")
	}

	return strings.Join(placeholders, " AND ")
}

// create executes passed query and args
func create(db *DB, query string, args ...interface{}) (string, error) {
	id, err := db.UUID()

	if err != nil {
		return "", err
	}

	args = append([]interface{}{id}, args...)

	result, err := db.Exec(query, args...)

	if err != nil {
		return "", err
	}

	if _, err = result.RowsAffected(); err != nil {
		return "", err
	}

	return args[0].(string), nil
}
