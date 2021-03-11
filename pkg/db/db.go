package db

import (
	"fmt"
	"time"

	"github.com/anthonyzero/go-quick-api/configs"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *Database

//Database gorm.DB
type Database struct {
	Default *gorm.DB
}

//Init 初始化db
func Init() error {
	cfg := configs.Get().MySQL

	//默认db配置
	defaultDbConf := cfg.Default
	defaultDb, err := dbConnect(defaultDbConf.User, defaultDbConf.Pass, defaultDbConf.Addr, defaultDbConf.Name)
	if err != nil {
		return err
	}
	//new
	db = &Database{
		Default: defaultDb,
	}
	return nil
}

//Close 关闭db
func Close() error {
	sqlDB, err := db.Default.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

//GetDefault 获取默认数据库DB
func GetDefault() *gorm.DB {
	return db.Default
}

func dbConnect(user, pass, addr, dbName string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		user,
		pass,
		addr,
		dbName,
		true,
		"Local")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		//Logger: logger.Default.LogMode(logger.Info), // 日志配置
	})

	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("[db connection failed] Database name: %s", dbName))
	}

	db.Set("gorm:table_options", "CHARSET=utf8mb4")

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// mysql基础配置
	cfg := configs.Get().MySQL.Base
	// 设置连接池 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConn)

	// 设置最大连接数 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConn)

	// 设置最大连接超时
	sqlDB.SetConnMaxLifetime(time.Minute * cfg.ConnMaxLifeTime)

	return db, nil
}
