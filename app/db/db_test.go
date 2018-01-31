package db

import "testing"

func testAssert(t *testing.T, param string, expected string, actual string) {
	if expected != actual {
		t.Errorf("model returned wrong %v: \ngot \"%v\" \nwant \"%v\"", param, actual, expected)
	}
}
