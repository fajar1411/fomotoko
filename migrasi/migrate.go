package migrasi

import (
	"test/domain/model"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Bank{})
	db.AutoMigrate(&model.TopUp{})
	db.AutoMigrate(&model.Wallet{})
	// db.AutoMigrate(&model.Transaction{})
}
