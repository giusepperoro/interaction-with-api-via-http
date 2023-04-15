package database

import (
	"context"
	"github.com/giusepperoro/interaction-with-api-via-http/internal/entity"
	"gorm.io/gorm"
)

type DataBase struct {
	Conn *gorm.DB
}

type DbManager interface {
	AddToDatabase(ctx context.Context, data entity.Result) (bool, error)
}
