package database

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"review/internal/models"
	"review/internal/pkg/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var mdb *gorm.DB

func GetMysql() *gorm.DB {
	return mdb
}

type CloseFunc func()

func InitializeMysql() CloseFunc {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Cfg.Mysql.Username,
		config.Cfg.Mysql.Password,
		config.Cfg.Mysql.Host,
		config.Cfg.Mysql.Port,
		config.Cfg.Mysql.Database,
	)
	gormConfig := &gorm.Config{}
	sqlDebug := false
	if sqlDebug {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	} else {
		gormConfig.Logger = logger.Default.LogMode(logger.Error)
	}
	var err error
	mdb, err = gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		slog.Error("connect mysql server error: ", slog.Any("err", err))
		os.Exit(1)
	}
	var sqlDB *sql.DB
	sqlDB, err = mdb.DB()
	if err != nil {
		slog.Error("get sql.DB error: ", slog.Any("err", err))
		os.Exit(1)
	}

	sqlDB.SetMaxOpenConns(config.Cfg.Mysql.ConnPool.MaxOpenConns)
	sqlDB.SetMaxIdleConns(config.Cfg.Mysql.ConnPool.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(config.Cfg.Mysql.ConnPool.ConnMaxLifetime))
	sqlDB.SetConnMaxIdleTime(time.Second * time.Duration(config.Cfg.Mysql.ConnPool.ConnMaxIdleTime))
	// 自动迁移数据库
	if true {
		err = mdb.AutoMigrate(
			&models.User{},
			&models.LoginLog{},
		)
		if err != nil {
			slog.Error("mysql table autoMigrate error: ", slog.Any("err", err))
			os.Exit(1)
		}
	}
	return func() {
		if err = sqlDB.Close(); err != nil {
			slog.Error("close mysql connection error", slog.Any("err", err))
		} else {
			slog.Debug("mysql connection closed success.")
		}
	}
}
