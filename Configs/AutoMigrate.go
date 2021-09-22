package Configs

import (
	"golang-final-project/Models/Login"
	"golang-final-project/Models/Users"
)

func Migrate() {
	DB.AutoMigrate(&Users.User{}, &Login.LoginDataUsers{})
}
