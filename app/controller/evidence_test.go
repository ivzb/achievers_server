package controller

import "testing"

func TestEvidencesAfter(t *testing.T) {
	run(t, evidencesAfterTests)
}

func TestEvidencesLatest(t *testing.T) {
	run(t, evidencesLatestTests)
}

func TestEvidenceSingle(t *testing.T) {
	run(t, evidenceSingleTests)
}

func TestEvidenceCreate(t *testing.T) {
	run(t, evidenceCreateTests)
}
