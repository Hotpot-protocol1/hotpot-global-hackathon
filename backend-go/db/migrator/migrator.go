package migrator

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/markbates/pkger"
	"github.com/pkg/errors"
	migrate "github.com/rubenv/sql-migrate"
)

// MigrateDir represents a direction in which to perform schema migrations.
type MigrateDir string

const (
	// MigrateUp causes migrations to be run in the "up" direction.
	MigrateUp MigrateDir = "up"
	// MigrateDown causes migrations to be run in the "down" direction.
	MigrateDown MigrateDir = "down"
)

type MigrationsLoader struct {
	source *migrate.HttpFileSystemMigrationSource
}

type Migrator func(*sql.DB, MigrateDir) (int, error)

func MigrateDB(direction string, dbClient *sqlx.DB, migratorFn Migrator) (int, error) {
	applied, err := migratorFn(dbClient.DB, MigrateDir(direction))

	return applied, errors.Wrap(err, "failed to apply migrations")
}

func NewMigrationsLoader() *MigrationsLoader {
	return &MigrationsLoader{}
}

func (l *MigrationsLoader) loadDir(dir string) {
	l.source = &migrate.HttpFileSystemMigrationSource{
		FileSystem: pkger.Dir(dir),
	}
}

// Migrate performs schema migration.  Migrations can occur in one of three
// ways:
//
// - up: migrations are performed from the currently installed version upwards.
// If count is 0, all unapplied migrations will be run.
//
// - down: migrations are performed from the current version downward. If count
// is 0, all applied migrations will be run in a downward direction.
//
// - redo: migrations are first ran downward `count` times, and then are ran
// upward back to the current version at the start of the process. If count is
// 0, a count of 1 will be assumed.
func (l *MigrationsLoader) Migrate(dbClient *sql.DB, dir MigrateDir) (int, error) {
	switch dir {
	case MigrateUp:
		return migrate.ExecMax(dbClient, "postgres", l.source, migrate.Up, 0)

	case MigrateDown:
		return migrate.ExecMax(dbClient, "postgres", l.source, migrate.Down, 0)

	default:
		return 0, errors.New("Invalid migration direction")
	}
}
