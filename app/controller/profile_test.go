package controller

import "testing"

func TestProfileMe(t *testing.T) {
	run(t, profileMeTests)
}

func TestProfileSingle(t *testing.T) {
	run(t, profileSingleTests)
}
