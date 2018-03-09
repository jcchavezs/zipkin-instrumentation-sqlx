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

type DB struct {
	db *sqlx.DB
	tr *zipkin.Tracer
}

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

func (db *DB) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	s, _ := db.tr.StartSpanFromContext(ctx, getNameFromQuery(query), zipkin.Tags(map[string]string{
		"sql.query": query,
	}))
	defer s.Finish()

	r, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		s.Tag(string(zipkin.TagError), err.Error())
	}

	return r, err
}

func (db *DB) QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	s, _ := db.tr.StartSpanFromContext(ctx, getNameFromQuery(query), zipkin.Tags(map[string]string{
		"sql.query": query,
	}))
	defer s.Finish()

	r, err := db.QueryxContext(ctx, query, args...)
	if err != nil {
		s.Tag(string(zipkin.TagError), err.Error())
	}

	return r, err
}

func (db *DB) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	s, _ := db.tr.StartSpanFromContext(ctx, getNameFromQuery(query), zipkin.Tags(map[string]string{
		"sql.query": query,
	}))
	defer s.Finish()

	r := db.QueryRowxContext(ctx, query, args...)
	if r.Err() != nil {
		s.Tag(string(zipkin.TagError), r.Err().Error())
	}

	return r
}

func (db *DB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	s, _ := db.tr.StartSpanFromContext(ctx, getNameFromQuery(query), zipkin.Tags(map[string]string{
		"sql.query": query,
	}))
	defer s.Finish()

	r, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		s.Tag(string(zipkin.TagError), err.Error())
	}

	return r, err
}

func (db *DB) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return db.PrepareContext(ctx, query)
}
