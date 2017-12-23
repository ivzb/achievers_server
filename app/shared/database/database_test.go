package database

import (
	"strconv"
	"testing"
)

func TestDSN_ValidMySQLInfo(t *testing.T) {
	mi := MySQLInfo{
		Username:  "root",
		Password:  "123",
		Name:      "achievers",
		Hostname:  "127.0.0.1",
		Port:      3306,
		Parameter: "?parseTime=true",
	}

	expected := mi.Username + ":" + mi.Password + "@tcp(" + mi.Hostname + ":" + strconv.Itoa(mi.Port) + ")/" + mi.Name + mi.Parameter

	actual := DSN(mi)

	if expected != actual {
		t.Fatalf("DSN returned wrong value: \nexpected %v, \nactual %v",
			expected, actual)
	}
}
