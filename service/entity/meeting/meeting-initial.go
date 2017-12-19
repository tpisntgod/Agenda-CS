package meeting

import (
	"fmt"

	"github.com/bilibiliChangKai/Agenda-CS/service/orm"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

//MeetingDB 数据库
var MeetingDB *xorm.Engine

func init() {
	err := orm.Mydb.Sync2(new(Meeting))
	fmt.Println("meeting database init")
	if err != nil {
		panic(err)
	}
	MeetingDB = orm.Mydb
}

//checkErr 输出错误
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
