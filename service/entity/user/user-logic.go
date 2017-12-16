package user

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/bilibiliChangKai/Agenda-CS/service/entity/mylog"
)

var userItemsFilePath = "src/github.com/bilibiliChangKai/Agenda-CS/service/orm/UserItems.json"
var currentUserFilePath = "src/github.com/bilibiliChangKai/Agenda-CS/service/orm/Current.txt"

func init() {
	// 初始化
	userItemsFilePath = filepath.Join(*mylog.GetGOPATH(), userItemsFilePath)
	currentUserFilePath = filepath.Join(*mylog.GetGOPATH(), currentUserFilePath)
}

// IsLogin : 判断当前有没有用户登录，并不是很必要
func IsLogin(name string) bool {
	return name != "" && service.FindByName(name) != nil
}

// RegisterUser : 注册用户，如果用户名一样，则返回err
func RegisterUser(name string, password string,
	email string, phoneNumber string) error {
	if service.FindByName(name).Name != "" {
		return errors.New("ERROR:The user has registered")
	}

	pitem := newItem(name, password, email, phoneNumber)
	service.Save(pitem)

	mylog.AddLog("", "RegisterUser", "", pitem.String())
	return nil
}

// LoginUser : 登录用户
// 如果用户名不存在，则返回err1
// 或者用户名密码不正确，则返回err2
func LoginUser(name string, password string, loginname string) (*Item, error) {
	if IsLogin(loginname) {
		return nil, errors.New("ERROR:Please logout at first")
	}
	pitem := service.FindByName(name)
	// 账号错误
	if pitem == nil {
		return nil, errors.New("ERROR:The user's name not exists")
	}

	// 密码错误
	if pitem.HashPassword != password {
		return nil, errors.New("ERROR:The user's password is wrong")
	}
	fmt.Println(pitem)

	// 成功登录
	mylog.AddLog(name, "LoginUser", "", "")
	return pitem, nil
}

// LogoutUser : 登出用户，如果当前没有用户登录，则返回err
func LogoutUser(loginname string) error {
	if !IsLogin(loginname) {
		return errors.New("ERROR:No login user")
	}

	mylog.AddLog(loginname, "LogoutUser", "", "")
	fmt.Println("Logout successfully!")
	return nil
}

// ListUsers : 返回所有用户信息
// 如果当前没有用户登录，返回err
func ListUsers(loginname string) ([]Item, error) {
	if !IsLogin(loginname) {
		return nil, errors.New("ERROR:No registered user")
	}

	returnItems := service.FindAll()
	mylog.AddLog(loginname, "ListUsers", "", "")
	return returnItems, nil
}

// DeleteUser : 删除当前登录用户，删除后当前登录用户置为nil
// 如果当前没有用户登录，返回err
func DeleteUser(loginname string) error {
	if !IsLogin(loginname) {
		return errors.New("ERROR:No registered user")
	}

	service.DeleteByName(loginname)
	mylog.AddLog(loginname, "DeleteUser", "", "")
	return nil
}

// to string
func (u Item) String() string {
	return "{Name:" + u.Name + "  Email:" + u.Email + "  Phone:" + u.PhoneNumber + "}"
}
