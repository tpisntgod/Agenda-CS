package user

import (
	"github.com/bilibiliChangKai/Agenda-CS/service/orm"
)

// ItemAtomicService 一个空类型
type ItemAtomicService struct{}

// service 空类型的指针，使用函数
var service = ItemAtomicService{}

func init() {
<<<<<<< HEAD
	err := orm.Mydb.Sync(new(Item))
=======
	err := orm.Mydb.Sync2(new(Item))
>>>>>>> f890abee758803ed374098a08edcc33a2abca2f7
	orm.CheckErr(err)
}

// Save 保存
func (*ItemAtomicService) Save(u *Item) error {
	_, err := orm.Mydb.Table("item").Insert(u)
	orm.CheckErr(err)
	return err
}

// FindAll 找到所有Item
func (*ItemAtomicService) FindAll() []Item {
	as := []Item{}
	err := orm.Mydb.Table("item").Desc("Name").Find(&as)
	orm.CheckErr(err)
	return as
}

// FindByName 通过主键Name查询数据
func (*ItemAtomicService) FindByName(name string) *Item {
	a := &Item{}
	_, err := orm.Mydb.Table("item").Id(name).Get(a)
	orm.CheckErr(err)
	return a
}

// DeleteByName 通过主键Name删除数据
func (*ItemAtomicService) DeleteByName(name string) {
	// 软删除
	_, err := orm.Mydb.Table("item").Id(name).Delete(&Item{})
	orm.CheckErr(err)

	// 真正删除
	_, err = orm.Mydb.Table("item").Id(name).Unscoped().Delete(&Item{})
	orm.CheckErr(err)
}
