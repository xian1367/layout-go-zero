package orm

import (
	"errors"
	"fmt"
	"github.com/xian1367/layout-go-zero/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

// Init 初始化数据库和 ORM
func Init() {
	var dbConfig gorm.Dialector
	switch config.Get().Database.Connection {
	case "mysql":
		// 构建 DSN 信息
		dbConfig = mysql.New(mysql.Config{
			DSN: fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
				config.Get().Database.Mysql.Username,
				config.Get().Database.Mysql.Password,
				config.Get().Database.Mysql.Host,
				config.Get().Database.Mysql.Port,
				config.Get().Database.Mysql.DBName,
				config.Get().Database.Mysql.Charset,
			),
		})
	case "postgres":
		// 初始化 sqlite
		postgres.New(postgres.Config{
			DSN: fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
				config.Get().Database.Postgres.Host,
				config.Get().Database.Postgres.Port,
				config.Get().Database.Postgres.Username,
				config.Get().Database.Postgres.Password,
				config.Get().Database.Postgres.DBName,
			),
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		})
	default:
		panic(errors.New("orm connection not supported"))
	}

	// 连接数据库，并设置 GORM 的日志模式
	Connect(dbConfig, NewGormLogger())
	// 设置最大连接数
	SqlDB.SetMaxOpenConns(config.Get().Database.Mysql.MaxOpenConnections)
	// 设置最大空闲连接数
	SqlDB.SetMaxIdleConns(config.Get().Database.Mysql.MaxIdleConnections)
	// 设置每个链接的过期时间
	SqlDB.SetConnMaxLifetime(time.Duration(config.Get().Database.Mysql.MaxLifeSeconds) * time.Second)
}
