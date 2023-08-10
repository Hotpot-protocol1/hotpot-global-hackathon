package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/logrusadapter"
	"github.com/sirupsen/logrus"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// DBHandler represent interface to implement for database handling
type DBHandler interface {
	DB() *sqlx.DB
	UserTickets() UserTickets
}

// DB wraps dbx interface.
type DB struct {
	db *sqlx.DB
}

// DB returns db client.
func (d *DB) DB() *sqlx.DB {
	return d.db
}

func New(connectionString string) DBHandler {
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	logger := logrus.New()
	logger.SetLevel(logrus.ErrorLevel)
	logger.SetReportCaller(true)

	db = sqldblogger.OpenDriver(connectionString, db.Driver(), logrusadapter.New(logger))

	return &DB{db: sqlx.NewDb(db, "pgx")}
}

func Close(db *sql.DB) {
	db.Close()
}
