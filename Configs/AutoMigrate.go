package Configs

import (
	"golang-final-project/Models"
)

func Migrate() {
	DB.AutoMigrate(&Models.User{})
}
