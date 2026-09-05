package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

func MigrateDB(db *sqlx.DB, dir string) error {
	migrations := migrate.FileMigrationSource{Dir: dir}

	total_migrated, err := migrate.Exec(db.DB, "postgres", migrations, migrate.Up)

	if err != nil {
		return err
	}
	fmt.Println("Total migrated", total_migrated)
	return nil
}
