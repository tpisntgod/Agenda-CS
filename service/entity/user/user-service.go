package user

import (
	"github.com/bilibiliChangKai/Agenda-CS/service/entity/mylog"
	"github.com/bilibiliChangKai/Agenda-CS/service/orm"
)

// ItemAtomicService 一个空类型
type ItemAtomicService struct{}

// service 空类型的指针，使用函数
var service = ItemAtomicService{}

func init() {
	err := orm.Mydb.Sync2(new(Item))
	checkErr(err)
}

// Save 保存
func (*ItemAtomicService) Save(u *Item) error {
	_, err := orm.Mydb.Table("item").Insert(u)
	checkErr(err)
	return err
}

// FindAll 找到所有Item
func (*ItemAtomicService) FindAll() []Item {
	as := []Item{}
	err := orm.Mydb.Table("item").Desc("Name").Find(&as)
	checkErr(err)
	return as
}

// FindByName 通过主键Name查询数据
func (*ItemAtomicService) FindByName(name string) *Item {
	a := &Item{}
	_, err := orm.Mydb.Table("item").Id(name).Get(a)
	checkErr(err)
	return a
}

// DeleteByName 通过主键Name删除数据
func (*ItemAtomicService) DeleteByName(name string) {
	// 软删除
	_, err := orm.Mydb.Table("item").Id(name).Delete(&Item{})
	checkErr(err)

	// 真正删除
	_, err = orm.Mydb.Table("item").Id(name).Unscoped().Delete(&Item{})
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		mylog.AddErr(err)
		panic(err)
	}
}
