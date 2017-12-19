package orm

import (
	"fmt"
	"path/filepath"

	"github.com/bilibiliChangKai/Agenda-CS/service/entity/mylog"
	"github.com/go-xorm/xorm"
	// 使用sqlite3数据库
	_ "github.com/mattn/go-sqlite3"
)

// Mydb 数据库指针
var Mydb *xorm.Engine

// 生成数据库，对数据库进行链接
func init() {
	// 链接sqlite3数据库
	//db, err := xorm.NewEngine("sqlite3", "./agenda-cs.db")
	databasePath := "src/github.com/bilibiliChangKai/Agenda-CS/service/agenda-cs.db"
	fmt.Println(databasePath)
	databasePath = filepath.Join(*mylog.GetGOPATH(), databasePath)
	db, err := xorm.NewEngine("sqlite3", databasePath)
	if err != nil {
		panic(err)
	}

	Mydb = db
}

// CheckErr panic错误
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
