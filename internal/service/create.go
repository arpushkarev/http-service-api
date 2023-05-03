package service

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	table      = "users"
	host       = "localhost"
	port       = "54321"
	dbUSer     = "zt-service-user"
	dbPassword = "zt-service-password"
	dbName     = "zt-service"
	sslMode    = "disable"
)

func () Create(ctx, context.Context, req) (int64, error) {
	dbDSN := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, dbUSer, dbPassword, dbName, sslMode,
	)

	db, err := sqlx.Open("pgx", dbDSN)
	if err != nil {
		return 0, err
	}
	defer db.Close()

}
