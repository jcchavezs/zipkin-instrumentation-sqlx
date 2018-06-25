package zipkinsqlx

import (
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/openzipkin/zipkin-go"
)

func assertExt(_ sqlx.Ext)               {}
func assertExtContext(_ sqlx.ExtContext) {}

func TestDBReturnsDriverName(t *testing.T) {
	driverName := "driver"
	tracer, _ := zipkin.NewTracer(nil)
	db := NewDb(sqlx.NewDb(nil, driverName), tracer)
	if want, have := driverName, db.DriverName(); want != have {
		t.Fatalf("unexpected driver name, expected %s, got %s", want, have)
	}

	assertExt(db)
	assertExtContext(db)
}

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
