package uuid

import (
	"crypto/rand"
	"fmt"
)

type UUIDer interface {
	Generate() (string, error)
}

type UUID struct {
}

func NewUUID() *UUID {
	return &UUID{}
}

func (uuid *UUID) Generate() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:]), nil
}
