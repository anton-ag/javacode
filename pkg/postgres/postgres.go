package postgres

import (
	"database/sql"
	"fmt"

	"github.com/anton-ag/javacode/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect(cfg *config.Config) (*sql.DB, error) {
	connString := "postgresql://" + cfg.Postgres.User + ":" + cfg.Postgres.Password + "@" + cfg.Postgres.Url + ":" + cfg.Postgres.Port + "/" + cfg.Postgres.Name
	db, err := sql.Open("pgx", connString)
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к БД: %w", err)
	}

	return db, nil
}

func InitTable(db *sql.DB) error {
	// TODO: move to file
	query := `CREATE TABLE IF NOT EXISTS wallet (
  id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  total integer DEFAULT 0
	);`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
