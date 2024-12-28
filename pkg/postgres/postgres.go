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

// TODO: Initialize empty database
