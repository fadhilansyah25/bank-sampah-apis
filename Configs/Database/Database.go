package Database

import (
	"fmt"

	"gorm.io/driver/postgres"
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

// // mysql
// func (dbConfig *DBConfig) DbURLMain() string {
// 	return fmt.Sprintf(
// 		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
// 		dbConfig.User,
// 		dbConfig.Password,
// 		dbConfig.Host,
// 		dbConfig.Port,
// 		dbConfig.DBName,
// 	)
// }

// func Connection(dsn DBConfig) {
// 	DB, err = gorm.Open(mysql.Open(dsn.DbURLMain()), &gorm.Config{})

// 	if err != nil {
// 		fmt.Println("Status:", err)
// 	}
// 	Migrate()
// }

// postgresSQL
func (dbConfig *DBConfig) DbURLMain() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Asia/Jakarta",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBName,
		dbConfig.Port,
	)
}

func Connection(dsn DBConfig) {
	DB, err = gorm.Open(postgres.Open(dsn.DbURLMain()), &gorm.Config{})

	if err != nil {
		fmt.Println("Status:", err)
	}
	Migrate()
}
