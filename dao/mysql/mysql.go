package mysql

import (
	"fmt"
	"gin-web/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 将每一步数据库操作封装成函数，待 logic 层根据业务需求调用

// sqlx 是对 database/sql 的封装，对于一些不熟悉的操作可以查看 sqlx 或者 database/sql 库的文档，例如模型的db标签等等
var db *sqlx.DB

func Init(cfg config.MysqlConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Dbname,
	)
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(cfg.MaxOpen)
	db.SetMaxIdleConns(cfg.MaxIdle)
	return
}

func Close() {
	_ = db.Close()
}
