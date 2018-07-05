// SQLx instrumented library with zipkin-go

package zipkinsqlx

import (
	"context"
	"database/sql"
	"regexp"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/reporter"
)

// DB is a wrapper around sqlx.DB with also required the zipkin tracer
type DB struct {
	db *sqlx.DB
	tr *zipkin.Tracer
}

// NewDb returns a new sqlx DB wrapper for a pre-existing *sqlx.DB.
func NewDb(db *sqlx.DB, tr *zipkin.Tracer) *DB {
	if tr == nil {
		tr, _ = zipkin.NewTracer(reporter.NewNoopReporter(), zipkin.WithNoopSpan(true))
	}

	return &DB{db, tr}
}

func getNameFromQuery(query string) string {
	re := regexp.MustCompile("[a-zA-Z]+")
	return strings.ToLower(re.FindString(query))
}

// QueryContext executes a query that returns rows, typically a SELECT.
// The args are for any placeholder parameters in the query.
func (db *DB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return db.db.QueryContext(context.Background(), query, args)
}

// QueryContext executes a query that returns rows, typically a SELECT.
// The args are for any placeholder parameters in the query.
func (db *DB) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	s, _ := db.tr.StartSpanFromContext(ctx, getNameFromQuery(query), zipkin.Tags(map[string]string{
		"sql.query": query,
	}))
	defer s.Finish()

	r, err := db.db.QueryContext(ctx, query, args...)
	if err != nil {
		s.Tag(string(zipkin.TagError), err.Error())
	}

	return r, err
}

// Queryx queries the database and returns an *sqlx.Rows.
// Any placeholder parameters are replaced with supplied args.
func (db *DB) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	return db.db.QueryxContext(context.Background(), query, args)
}

// QueryxContext queries the database and returns an *sqlx.Rows.
// Any placeholder parameters are replaced with supplied args.
func (db *DB) QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	s, _ := db.tr.StartSpanFromContext(ctx, getNameFromQuery(query), zipkin.Tags(map[string]string{
		"sql.query": query,
	}))
	defer s.Finish()

	r, err := db.db.QueryxContext(ctx, query, args...)
	if err != nil {
		s.Tag(string(zipkin.TagError), err.Error())
	}

	return r, err
}

// QueryRowx queries the database and returns an *sqlx.Row.
// Any placeholder parameters are replaced with supplied args.
func (db *DB) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	return db.db.QueryRowxContext(context.Background(), query, args)
}

// QueryRowxContext queries the database and returns an *sqlx.Row.
// Any placeholder parameters are replaced with supplied args.
func (db *DB) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	s, _ := db.tr.StartSpanFromContext(ctx, getNameFromQuery(query), zipkin.Tags(map[string]string{
		"sql.query": query,
	}))
	defer s.Finish()

	r := db.db.QueryRowxContext(ctx, query, args...)
	if r.Err() != nil {
		s.Tag(string(zipkin.TagError), r.Err().Error())
	}

	return r
}

// Exec execs a statement in the database and returns an sql.Result.
// Any placeholder parameters are replaced with supplied args.
func (db *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.db.ExecContext(context.Background(), query, args)
}

// ExecContext execs a statement in the database and returns an sql.Result.
// Any placeholder parameters are replaced with supplied args.
func (db *DB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	s, _ := db.tr.StartSpanFromContext(ctx, getNameFromQuery(query), zipkin.Tags(map[string]string{
		"sql.query": query,
	}))
	defer s.Finish()

	r, err := db.db.ExecContext(ctx, query, args...)
	if err != nil {
		s.Tag(string(zipkin.TagError), err.Error())
	}

	if affectedRows, err := r.RowsAffected(); err != nil {
		s.Tag("db.affected_rows", string(affectedRows))
	}

	return r, err
}

// PreparexContext prepares a statement.
//
// The provided context is used for the preparation of the statement, not for
// the execution of the statement.
func (db *DB) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return db.db.PrepareContext(ctx, query)
}

func (db *DB) DriverName() string {
	return db.db.DriverName()
}

func (db *DB) Rebind(query string) string {
	return db.db.Rebind(query)
}

func (db *DB) BindNamed(query string, arg interface{}) (string, []interface{}, error) {
	return db.db.BindNamed(query, arg)
}
