package migration

import (
	"gorm.io/gorm"
	"ms-workspace/ms-user/internal/v1/repository/model"
)

func CreateTable(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.ActiveCode{})
}
