package db

type Involvementer interface {
	Exists(id string) (bool, error)
}

type Involvement struct {
	db *DB
}

func (db *DB) Involvement() Involvementer {
	return &Involvement{db}
}

func (ctx *Involvement) Exists(id string) (bool, error) {
	return exists(ctx.db, "involvement", "id", id)
}
