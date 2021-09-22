package Configs

import (
	"golang-final-project/Models/BankSampah"
	"golang-final-project/Models/Users"
)

func Migrate() {
	DB.AutoMigrate(&Users.User{}, &Users.LoginDataUsers{}, &BankSampah.BankSampah{})
}
