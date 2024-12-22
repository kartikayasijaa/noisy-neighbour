package models

import "noisy_neighbour/common/database/base_model"

type UserRecord struct {
	base_model.BaseModel
	SiteID   uint   `json:"site_id"`
	Site     Site   `json:"site" gorm:"foreignKey:SiteID;constraint:OnDelete:CASCADE"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	IsActive bool   `json:"is_active" gorm:"default:true"`
}
