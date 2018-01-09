package controller

import "testing"

func TestUserAuth(t *testing.T) {
	run(t, userAuthTests)
}

func TestUserCreate(t *testing.T) {
	run(t, userCreateTests)
}
