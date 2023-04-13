package database

import (
	"context"
	"github.com/giusepperoro/interaction-with-api-via-http/internal/config"
	"github.com/giusepperoro/interaction-with-api-via-http/internal/entity"
	"github.com/jackc/pgx/v4/pgxpool"
)

func New(ctx context.Context, cfg config.ServiceConfiguration) (*DataBase, error) {
	connection, err := pgxpool.Connect(ctx, cfg.PostgresConnectUrl)
	if err != nil {
		return nil, err
	}
	err = connection.Ping(ctx)
	if err != nil {
		return nil, err
	}
	return &DataBase{Conn: connection}, nil
}

func (d *DataBase) AddToDatabase(ctx context.Context, data entity.Result) (bool, error) {
	return false, nil
}
