package db

import (
	"context"
	"fmt"
	"time"

	"github.com/yuki-inoue-eng/authenticator/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB interface {
	NewDB(ctx context.Context, config configs.Configs) error
}

type MySQL struct {
	Con  *gorm.DB
	Info Info
}

type Info struct {
	DBName    string
}

var ConPool = &MySQL{}

func NewDB(ctx context.Context, config configs.Configs) error {
	cfgs := config.Get()
	ConPool.Info = Info{
		DBName:    cfgs.Database.Name,
	}
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		cfgs.Database.UserName,
		cfgs.Database.UserPassword,
		cfgs.Database.HostName,
		cfgs.Database.Port,
		cfgs.Database.Name,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	ConPool.Con = db.WithContext(ctx)
	mysqlDB, err := db.DB()
	if err != nil {
		return err
	}
	err = mysqlDB.Ping()
	if err != nil {
		panic(err)
	}
	mysqlDB.SetConnMaxIdleTime(24 * time.Hour)
	mysqlDB.SetMaxOpenConns(cfgs.Database.MaxOpenCon)
	mysqlDB.SetMaxIdleConns(cfgs.Database.MaxIdleCon)
	return nil
}