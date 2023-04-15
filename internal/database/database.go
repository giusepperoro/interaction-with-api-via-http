package database

import (
	"context"
	"github.com/giusepperoro/interaction-with-api-via-http/internal/config"
	"github.com/giusepperoro/interaction-with-api-via-http/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(cfg config.ServiceConfiguration) (*DataBase, error) {
	connection, err := gorm.Open(postgres.Open(cfg.PostgresConnectUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &DataBase{Conn: connection}, nil
}

func (d *DataBase) AddToDatabase(ctx context.Context, data entity.Result) (bool, error) {

	return false, nil
}
