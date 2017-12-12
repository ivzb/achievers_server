package model

type Env struct {
	DB    DBSource
	Store map[string]string
	Token TokenSource
}
