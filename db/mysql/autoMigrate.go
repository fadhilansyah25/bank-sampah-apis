package mysql

import "gorm.io/gorm"

func Migrate(DB *gorm.DB) {
	DB.AutoMigrate()
}
