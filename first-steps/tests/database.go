package database

import "database/sql"

func Connect(dsn string) (*sql.DB, error) {
	return sql.Open("postgres", dsn)
}
