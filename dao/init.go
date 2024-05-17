package dao

import (
	"fmt"

	"alex.tse/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dbInstance *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DataBase string
}

func GetDBInstance() *gorm.DB {
	return dbInstance
}

func InitDB() error {
	var err error
	dbConfig := &gorm.Config{
		SkipDefaultTransaction: true, // 禁用默认事务
		Logger:                 logger.Default.LogMode(logger.Error),
	}
	conf := &DBConfig{
		Host:     "gz-cdb-6zazw99f.sql.tencentcdb.com",
		Port:     28142,
		Username: "root",
		Password: "Xiewenjie0922!",
		DataBase: "hupu",
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Username, conf.Password, conf.Host, conf.Port, conf.DataBase)
	dbInstance, err = gorm.Open(mysql.Open(dsn), dbConfig)
	if err != nil {
		return fmt.Errorf("failed to connect to database, error: %w", err)
	}

	if err = dbInstance.AutoMigrate(&model.RaceType{}); err != nil {
		return fmt.Errorf("failed to auto migrate: %v", err)
	}

	return nil
}
