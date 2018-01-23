package database

import (
	"strconv"
	"testing"
)

// postgres://username:password@localhost:port/db_name?sslmode=disable
func TestDSN_ValidMySQLInfo(t *testing.T) {
	pi := PostgreInfo{
		Username:  "Admin",
		Password:  "",
		Name:      "achievers",
		Hostname:  "127.0.0.1",
		Port:      5432,
		Parameter: " sslmode=disable",
	}

	expected := "postgres://" + pi.Username + ":" + pi.Password + "@" + pi.Hostname + ":" + strconv.Itoa(pi.Port) + "/" + pi.Name + "?" + pi.Parameter

	actual := DSN(pi)

	if expected != actual {
		t.Fatalf("DSN returned wrong value: \nexpected %v, \nactual %v",
			expected, actual)
	}
}
