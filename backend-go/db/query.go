package db

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// Queryable type is used as handle for functions that could be called with both sql(x).DB and sql(x).Tx depending on use-case
type Queryable interface {
	PrepareNamed(query string) (*sqlx.NamedStmt, error)
	NamedExec(query string, arg interface{}) (sql.Result, error)
	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)
	QueryRowx(query string, args ...interface{}) *sqlx.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
}
