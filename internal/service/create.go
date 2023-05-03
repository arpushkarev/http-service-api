package service

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/arpushkarev/http-service-api/cmd"
	_ "github.com/jackc/pgx/stdlib"
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

func () Create(ctx, context.Context, req cmd.Person) (int64, error) {
	dbDSN := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, dbUSer, dbPassword, dbName, sslMode,
	)

	db, err := sqlx.Open("pgx", dbDSN)
	if err != nil {
		return 0, err
	}
	defer db.Close()

	builder := sq.Insert(table).
	PlaceholderFormat(sq.Dollar).
	Columns("name", "age").
	Values(req.Name(), req.Age()).
	Suffix("returning id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	row, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	defer row.Close()

	row.Next()
	var id int64
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
