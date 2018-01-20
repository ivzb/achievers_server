package db

type Profiler interface {
	ProfileExists(id string) (bool, error)
	ProfileSingle(id string) (*Profile, error)
	ProfileByUserID(userID string) (*Profile, error)
	ProfileCreate(profile *Profile, userID string) (string, error)
}

type Profile struct {
	db *DB
}

func NewProfile(db *DB) *Profile {
	return &Profile{db: db}
}
