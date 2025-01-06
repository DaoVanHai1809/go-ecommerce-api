package initialize

import (
	"fmt"
	"go-ecommerce-backend-api/global"
	"go-ecommerce-backend-api/internal/po"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMySql() {
	var m = global.Config.Mysql
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.DbName)
  	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	checkErrorPanic(err, "InitMysql initialization error")
	global.Logger.Info("Initializing mysql successfully")
	global.Mdb = db

	// set pool
	SetPool()
	migrateTables()
}

func SetPool() {
	var m = global.Config.Mysql
	sqlDb, err := global.Mdb.DB()
	if err != nil {
		fmt.Printf("mysql error: %s::", err)
	}
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifeTime))
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
}

func migrateTables() {
	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)
	if err != nil {
		fmt.Println("migrating table error ", err)
	}
}