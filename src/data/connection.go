package data

import (
	"github.com/A-Siam/go_auth/src/data/entities"
	"github.com/A-Siam/go_auth/src/services/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var host = utils.GetEnv("PGSQL_HOST")
var user = utils.GetEnv("PGSQL_USERNAME")
var password = utils.GetEnv("PGSQL_PASSWORD")
var dbName = utils.GetEnv("PGSQL_DB_NAME")
var port = utils.GetEnv("PGSQL_PORT")

var dsn = "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbName + " port=" + port + " sslmode=disable"
var db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

func GetDB() *gorm.DB {
	db.AutoMigrate(&entities.User{})
	return db
}
