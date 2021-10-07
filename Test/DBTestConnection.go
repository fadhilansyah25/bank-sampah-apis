package Test

import (
	"fmt"
	"golang-final-project/Models/BankSampah"
	"golang-final-project/Models/Transaction"
	"golang-final-project/Models/UserLogins"
	"golang-final-project/Models/Users"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
}

// mysql
func (dbConfig *DBConfig) DbURLMain() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

func Connection(dsn DBConfig) {
	DB, err = gorm.Open(mysql.Open(dsn.DbURLMain()), &gorm.Config{})

	if err != nil {
		fmt.Println("Status:", err)
	}
	Migrate()
}

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

func SetUpTestDB() *gorm.DB {
	Connection(DBConfig{
		Host:     "localhost",
		User:     "root",
		Password: "",
		Port:     "3306",
		DBName:   "go_bank_sampah_test",
	})
	return DB
}
