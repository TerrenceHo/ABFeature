package stores

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Store is a postgres store
type Store interface {
	migrate()
}

// CreateTables creates tables for all stores
func CreateTables(stores ...Store) {
	for _, s := range stores {
		s.migrate()
	}
}

// TODO: Actually test SQLite engine, and add a SQLite Driver
func NewConnection(engine, user, password, dbname, port, host string) *sqlx.DB {
	var dbConnection string
	switch engine {
	case "postgres":
		dbConnection = connectPostgres(user, password, dbname, port, host)
	case "sqlite":
		dbConnection = connectSqlite()
	}

	db, err := sqlx.Connect(engine, dbConnection)
	if err != nil {
		fmt.Println("Error connecting to database", err)
		os.Exit(1)
	}

	return db
}

func connectPostgres(user, password, dbname, port, host string) string {
	return fmt.Sprintf("user=%s password=%s dbname=%s port=%s host=%s sslmode=disable",
		user, password, dbname, port, host)
}

func connectSqlite() string {
	return ":memory:"
}
