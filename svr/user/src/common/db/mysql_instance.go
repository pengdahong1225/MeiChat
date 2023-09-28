package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"user/src/common/initializer"
)

/*
 *	DB是一个数据库（操作）句柄，代表一个具有零到多个底层连接的连接池。它可以安全的被多个go程同时使用。
 *	Open返回的DB可以安全的被多个go程同时使用，并会维护自身的闲置连接池。这样一来，Open函数只需调用一次。很少需要关闭DB。
 */

func SqlInstance() *sql.DB {
	sqlInfo := initializer.DBInfoInstance.Sql_
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", sqlInfo.User, sqlInfo.Pwd, sqlInfo.Host, sqlInfo.Port, sqlInfo.DB)
	SqlInstance, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	SqlInstance.SetConnMaxLifetime(time.Minute * 10)
	SqlInstance.SetMaxOpenConns(10)
	SqlInstance.SetMaxIdleConns(10)

	// 验证连接
	if err := SqlInstance.Ping(); err != nil {
		panic(err)
	}
	return SqlInstance
}
