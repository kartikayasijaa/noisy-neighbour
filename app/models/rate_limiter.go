package models

import (
	"noisy_neighbour/common/database/base_model"
	"time"
)

type RateLimiter struct {
	base_model.BaseModel
	SiteID    uint      `json:"site_id" gorm:"uniqueIndex:idx_id_siteid"`
	Site      Site      `json:"site" gorm:"foreignKey:SiteID;constraint:OnDelete:CASCADE"`
	Rate      int       `json:"rate"`
	Capacity  int       `json:"capacity"`
	Tokens    int       `json:"tokens"`
	LastCheck time.Time `json:"last_check"`
}
