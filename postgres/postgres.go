package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// Connect connects to the Postgres database. This returns the connected database, or an error if there were any issues
// connecting to the database.
func Connect() (*sql.DB, error) {
	connectionString, err := getConnectionString()
	if err != nil {
		return nil, err
	}

	return sql.Open("pgx", connectionString)
}

func getConnectionString() (string, error) {
	user := os.Getenv("DB_USER")
	if user == "" {
		return "", errors.New("database user is not set")
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		return "", errors.New("database password is not set")
	}

	host := os.Getenv("DB_HOST")
	if host == "" {
		return "", errors.New("database host is not set")
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		return "", errors.New("database port is not set")
	}

	databaseName := os.Getenv("DB_NAME")
	if databaseName == "" {
		return "", errors.New("database name is not set")
	}

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, databaseName), nil
}
