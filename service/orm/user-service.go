package orm

import (
	"github.com/bilibiliChangKai/Agenda-CS/service/entity/user"
)

//UserItemAtomicService 一个空类型
type UserItemAtomicService struct{}

//UserItemService 空类型的指针，使用函数
var UserItemService = UserItemAtomicService{}

// Save 保存
func (*UserItemAtomicService) Save(u *user.Item) error {
	_, err := mydb.Insert(u)
	checkErr(err)
	return err
}

// FindAll 找到所有Item
func (*UserItemAtomicService) FindAll() []user.Item {
	as := []user.Item{}
	err := mydb.Desc("Name").Find(&as)
	checkErr(err)
	return as
}

// FindByName 通过主键Name查询数据
func (*UserItemAtomicService) FindByName(name string) *user.Item {
	a := &user.Item{}
	_, err := mydb.Id(name).Get(a)
	checkErr(err)
	return a
}

// DeleteByName 通过主键Name删除数据
func (*UserItemAtomicService) DeleteByName(name string) error {
	// 软删除
	mydb.Id(name).Delete(&user.Item{})
	// 真正删除
	mydb.Id(name).Unscoped().Delete(&user.Item{})
	return nil
}