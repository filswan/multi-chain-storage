package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"payment-bridge/off-chain/config"
	"payment-bridge/off-chain/logs"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

// Opening a database and save the reference to `Database` struct.
func Init() *gorm.DB {
	dbSource := config.GetConfig().Database.DbUsername + ":" + config.GetConfig().Database.DbPwd + "@tcp(" + config.GetConfig().Database.DbHost + ":" + config.GetConfig().Database.DbPort + ")/" + config.GetConfig().Database.DbSchemaName + "?" + config.GetConfig().Database.DbArgs
	db, err := gorm.Open("mysql", dbSource)
	if err != nil {
		logs.GetLogger().Fatal("db err: ", err)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	//db.LogMode(true)
	db.LogMode(true)
	DB = db
	return DB
}

// Using this function to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}

func SaveOne(data interface{}) error {
	db := GetDB()
	err := db.Save(data).Error
	return err
}
