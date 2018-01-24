package controller

import "testing"

func TestFileSingle(t *testing.T) {
	for _, test := range fileSingleTests {
		rec := constructRequest(t, test)

		assertCoreResponse(t, rec, test)
	}
}

func TestFileCreate(t *testing.T) {
	run(t, fileCreateTests)
}
