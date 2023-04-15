package handlers

import (
	"github.com/giusepperoro/interaction-with-api-via-http/internal/config"
	"github.com/giusepperoro/interaction-with-api-via-http/internal/entity"
	"gorm.io/gorm"
)

type DataRequest struct {
	gorm.Model
	results []entity.Result `json:"results"`
}

type DataSavedResponse struct {
	Status string `json:"status"`
}

type addToDatabaseHandler struct {
	cfg config.ServiceConfiguration
}

func NewAddToDatabaseHandler(cfg config.ServiceConfiguration) *addToDatabaseHandler {
	return &addToDatabaseHandler{
		cfg: cfg,
	}
}
