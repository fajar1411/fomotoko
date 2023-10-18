package migrasi

import (
	"test/domain/model"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Admin{})
	db.AutoMigrate(&model.Barang{})
	db.AutoMigrate(&model.OrderBarang{})
}
