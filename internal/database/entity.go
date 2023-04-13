package database

import (
	"context"
	"github.com/giusepperoro/interaction-with-api-via-http/internal/entity"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DataBase struct {
	Conn *pgxpool.Pool
}

type DbManager interface {
	AddToDatabase(ctx context.Context, data entity.Result) (bool, error)
}
