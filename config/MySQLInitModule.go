package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"pood/v2/app/models"
)

var Db *gorm.DB
var err error

func initDatabase() bool {

	mysqlUserName, _ := os.LookupEnv("MYSQL_USER")
	mysqlPassword, _ := os.LookupEnv("MYSQL_PASSWORD")
	mysqlHost, _ := os.LookupEnv("MYSQL_HOST")
	mysqlPort, _ := os.LookupEnv("MYSQL_PORT")
	mysqlDbName, _ := os.LookupEnv("MYSQL_DB_NAME")
	//dbDriver, _ := os.LookupEnv("DB_DRIVER")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", mysqlUserName, mysqlPassword, mysqlHost, mysqlPort, mysqlDbName)
	//dsn := fmt.Sprintf("%s://%s:%s@%s:%s/%s", dbDriver, mysqlUserName, mysqlPassword, mysqlHost, mysqlPort, mysqlDbName)
	log.Println(dsn)

	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	Db = Db.Debug()

	if err != nil {
		fmt.Println("DB initialization failed")
		panic(err)
		return false
	}

	return true
}

func initMigrations() bool {
	err := Db.AutoMigrate(
		models.User{},
		models.Action{},
		models.UserAction{},
		models.Log{},
		models.Token{},
		models.TypeInfo{},
		models.SubTypeInfo{},
		models.Unit{},
		models.File{},
	)

	if err != nil {
		fmt.Println("Migration failed")
		panic(err)
		return false
	}

	return true
}
