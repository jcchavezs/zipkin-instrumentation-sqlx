package zipkinsqlx

import "testing"

func TestGetNameFromQuery(t *testing.T) {
	queryProvider := [][]string{
		{"delete FROM", "delete"},
		{" SELECT * FROM", "select"},
		{`

	   UPDATE users SET
`, "update"},
	}

	for _, testData := range queryProvider {
		if want, have := testData[1], getNameFromQuery(testData[0]); want != have {
			t.Errorf("unexpected name, expected %s, got %s", want, have)
		}
	}
}

