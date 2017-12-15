package entity

import (
	"github.com/go-xorm/xorm"
	// 使用sqlite3数据库
	_ "github.com/mattn/go-sqlite3"
)

// 数据库指针
var mydb *xorm.Engine

// 生成数据库，对数据库进行链接
func init() {
	// 链接sqlite3数据库
	db, err := xorm.NewEngine("sqlite3", "./agenda-cs.db")

	if err != nil {
		panic(err)
	}

	// err = db.Sync(new(UserItem))
	// if err != nil {
	// 	panic(err)
	// }

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
