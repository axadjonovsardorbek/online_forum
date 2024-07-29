package postgresql

import (
	"auth-service/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDB(cf *config.Config) (*sql.DB, error) {
	dbConn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cf.DB_USER, cf.DB_PASSWORD, cf.DB_HOST, cf.DB_PORT, cf.DB_NAME)
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
