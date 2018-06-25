// Package zipkin-instrumentation-sqlx provides a traced wrapper for jmoiron/sqlx.
//
// None of the underlying jmoiron/sqlx methods are changed in functionality.
// Instead, they are wrapped with zipkin traces but output will be never be changed
// or affected.
//
package zipkinsqlx
