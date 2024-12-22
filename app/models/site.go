package models

import (
	"noisy_neighbour/common/database/base_model"
)

type Choice int

const (
	Low    Choice = 1
	Medium Choice = 2
	High   Choice = 3
)

type Site struct {
	base_model.BaseModel
	Name             string `json:"title" validate:"required"`
	Description      string `json:"description"`
	Domain           string `json:"domain"`
	URL              string `json:"url"`
	RecordCapacities Choice `json:"record_capacities"`
}
