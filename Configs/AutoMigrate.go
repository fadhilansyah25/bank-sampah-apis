package Configs

import (
	"golang-final-project/Models/BankSampah"
	"golang-final-project/Models/Transaction"
	"golang-final-project/Models/UserLogins"
	"golang-final-project/Models/Users"
)

func Migrate() {
	DB.AutoMigrate(
		&Users.User{},
		&UserLogins.LoginDataUsers{},
		&BankSampah.BankSampah{},
		&BankSampah.OperatorSampah{},
		&Transaction.JenisSampah{},
		&Transaction.Transaction{},
		&Transaction.DetailTransaction{},
	)
}
