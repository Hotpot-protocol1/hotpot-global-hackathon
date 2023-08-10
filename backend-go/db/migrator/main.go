package migrator

const (
	// MigrationsDir is folder where migrations are store
	MigrationsDir = "/db/migrator/migrations"
)

var (
	Migrations *MigrationsLoader
)

func init() {
	Migrations = NewMigrationsLoader()

	Migrations.loadDir(MigrationsDir)
}
