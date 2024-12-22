package helpers

import (
	"noisy_neighbour/app/models"
)

func GetAllModels() []interface{} {
	return []interface{}{
		&models.Site{},
		&models.UserRecord{},
		&models.RateLimiter{},
	}
}
