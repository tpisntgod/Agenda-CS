package orm

import (
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

var mydb *xorm.Engine

// 生成数据库，对数据库进行链接
func InitMydb(name string, password string, port string, dname string) {
	// 链接sqlite3数据库
	db, err := xorm.NewEngine("sqlite3", "./agenda-cs.db")

	if err != nil {
		panic(err)
	}

	// 同步注册表
	//err = db.Sync(new(UserInfo))
	if err != nil {
		panic(err)
	}

	mydb = db
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
