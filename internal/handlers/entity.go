package handlers

import "github.com/giusepperoro/interaction-with-api-via-http/internal/entity"

type dataRequest struct {
	results []entity.Result
}

type dataSavedResponse struct {
	Status string `json:"status"`
}