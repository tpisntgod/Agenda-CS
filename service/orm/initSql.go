package orm

import (
	"github.com/go-xorm/xorm"
	// 使用sqlite3数据库
	_ "github.com/mattn/go-sqlite3"

	"github.com/bilibiliChangKai/Agenda-CS/service/entity/user"
)

var mydb *xorm.Engine

// InitMydb 生成数据库，对数据库进行链接
func InitMydb(name string, password string, port string, dname string) {
	// 链接sqlite3数据库
	db, err := xorm.NewEngine("sqlite3", "./agenda-cs.db")

	if err != nil {
		panic(err)
	}

	// 同步user，meeting注册表
	err = db.Sync(new(user.Item))
	if err != nil {
		panic(err)
	}
	// err = db.Sync(new(meeting.))
	// if err != nil {
	// 	panic(err)
	// }

	mydb = db
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
