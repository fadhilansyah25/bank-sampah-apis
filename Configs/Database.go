package Configs

import (
	"fmt"

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

func (dbConfig *DBConfig) DbURL() string {
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
	DB, err = gorm.Open(mysql.Open(dsn.DbURL()), &gorm.Config{})

	if err != nil {
		fmt.Println("Status:", err)
	}
	Migrate()
}
