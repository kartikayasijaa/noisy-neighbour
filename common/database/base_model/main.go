package base_model

import (
	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model
}

type BaseManager struct {
	DB *gorm.DB
}

func NewBaseManager(db *gorm.DB) *BaseManager {
	return &BaseManager{DB: db}
}

func (bm *BaseManager) Save(value interface{}) error {
	return bm.DB.Save(value).Error
}

func (bm *BaseManager) Update(value interface{}) error {
	return bm.DB.Save(value).Error
}

func (bm *BaseManager) Delete(value interface{}) error {
	return bm.DB.Delete(value).Error
}

func (bm *BaseManager) Find(dest interface{}, id interface{}) error {
	return bm.DB.First(dest, id).Error
}
func (bm *BaseManager) Create(value interface{}) error {
	return bm.DB.Create(value).Error
}
