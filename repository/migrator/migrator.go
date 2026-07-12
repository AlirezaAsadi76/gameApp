package migrator

import (
	"database/sql"
	"fmt"
	"gameApp/repository/mysql"

	migrate "github.com/rubenv/sql-migrate"
)

type Migrator struct {
	dialect    string
	config     mysql.Config
	migrations *migrate.FileMigrationSource
}

func New(config mysql.Config) Migrator {

	migrations := &migrate.FileMigrationSource{
		Dir: "./repository/mysql/migrations",
	}

	return Migrator{
		dialect:    "mysql",
		config:     config,
		migrations: migrations,
	}
}

func (m Migrator) Up() {

	db, err := sql.Open(m.dialect, fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true",
		m.config.Username, m.config.Password, m.config.Host, m.config.Port, m.config.Database))
	if err != nil {
		panic(fmt.Errorf("can't open database: %v", err))
	}

	n, err := migrate.Exec(db, m.dialect, m.migrations, migrate.Up)
	if err != nil {
		panic(fmt.Errorf("can't Apply migration: %v", err))
	}
	fmt.Printf("Applied %d migrations!\n", n)
}

func (m Migrator) Down() {

	db, err := sql.Open(m.dialect, fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true",
		m.config.Username, m.config.Password, m.config.Host, m.config.Port, m.config.Database))
	if err != nil {
		panic(fmt.Errorf("can't open database: %v", err))
	}

	n, err := migrate.Exec(db, m.dialect, m.migrations, migrate.Down)
	if err != nil {
		panic(fmt.Errorf("can't rollback migration: %v", err))
	}
	fmt.Printf("rolled back %d migrations!\n", n)
}
