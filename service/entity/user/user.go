package user

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/tpisntgod/Agenda/service/entity/mylog"
)

var userItemsFilePath = "src/github.com/tpisntgod/Agenda/service/orm/UserItems.json"
var currentUserFilePath = "src/github.com/tpisntgod/Agenda/service/orm/Current.txt"

// Item 用户信息
type Item struct {
	// 用户名字，是唯一主键
	Name string `xorm:"pk" json:",omitempty"`
	// hash过的密码
	HashPassword string `json:"Password"`
	// 注册用的邮箱
	Email string `json:",omitempty"`
	// 注册用的电话号码
	PhoneNumber string `json:"Phone,omitempty"`
}

func init() {
	// 初始化
	userItemsFilePath = filepath.Join(*mylog.GetGOPATH(), userItemsFilePath)
	currentUserFilePath = filepath.Join(*mylog.GetGOPATH(), currentUserFilePath)
	userItems = make(map[string](Item))
	CurrentUser = nil
	readJSON()
}

// SetCurrentUser 新加入的函数，通过cookie字段，更新CurrentUser
func SetCurrentUser(name string, err error) {
	// cookie为空
	if err != nil {
		CurrentUser = nil
		return
	}

	// CurrentUser未改变
	if CurrentUser != nil && CurrentUser.Name == name {
		return
	}

	// CurrentUser改变，则更新
	tempUser, ok := userItems[name]
	if !ok {
		CurrentUser = nil
	} else {
		CurrentUser = &tempUser
	}
}

// 新建一个Item，并返回指针
func newUser(name string, password string,
	email string, phoneNumber string) *Item {
	newItem := new(Item)
	newItem.Name = name
	newItem.HashPassword = hashFunc(password)
	newItem.Email = email
	newItem.PhoneNumber = phoneNumber
	return newItem
}

// 用于密码hash的函数
func hashFunc(hashString string) string {
	// 进行md5加密
	h := md5.New()
	h.Write([]byte(hashString))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

// 储存user的map集合
var userItems map[string](Item)

// CurrentUser : currentUser是当前User，如果没有登录为nil
var CurrentUser *Item

// IsLogin : 判断当前有没有用户登录，并不是很必要
func IsLogin() bool {
	return CurrentUser != nil
}

// RegisterUser : 注册用户，如果用户名一样，则返回err
func RegisterUser(name string, password string,
	email string, phoneNumber string) error {
	_, ok := userItems[name]
	if ok {
		return errors.New("ERROR:The user has registered")
	}
	userItems[name] = *newUser(name, password, email, phoneNumber)

	writeJSON()
	mylog.AddLog("", "RegisterUser", "", userItems[name].String())
	return nil
}

// LoginUser : 登录用户
// 如果用户名不存在，则返回err1
// 或者用户名密码不正确，则返回err2
func LoginUser(name string, password string) error {
	if IsLogin() {
		return errors.New("ERROR:Please logout at first")
	}
	tempUser, nameok := userItems[name]
	// 账号错误
	if !nameok {
		return errors.New("ERROR:The user's name not exists")
	}

	passwordok := tempUser.HashPassword == hashFunc(password)
	// 密码错误
	if !passwordok {
		return errors.New("ERROR:The user's password is wrong")
	}

	// 成功登录
	CurrentUser = &tempUser
	writeJSON()
	mylog.AddLog(GetLogonUsername(), "LoginUser", "", "")
	fmt.Println("Hi " + name + "!")
	return nil
}

// LogoutUser : 登出用户，如果当前没有用户登录，则返回err
func LogoutUser() error {
	if !IsLogin() {
		return errors.New("ERROR:No login user")
	}

	CurrentUser = nil
	writeJSON()
	mylog.AddLog(GetLogonUsername(), "LogoutUser", "", "")
	fmt.Println("Logout successfully!")
	return nil
}

// ListUsers : 返回所有用户信息
// 如果当前没有用户登录，返回err
func ListUsers() ([]Item, error) {
	if !IsLogin() {
		return nil, errors.New("ERROR:No registered user")
	}

	var returnItems []Item

	// 依次添加新元素到数组
	for _, user := range userItems {
		returnItems = append(returnItems, user)
	}

	mylog.AddLog(GetLogonUsername(), "ListUsers", "", "")

	return returnItems, nil
}

// DeleteUser : 删除当前登录用户，删除后当前登录用户置为nil
// 如果当前没有用户登录，返回err
func DeleteUser() error {
	if !IsLogin() {
		return errors.New("ERROR:No registered user")
	}

	delete(userItems, CurrentUser.Name)

	mylog.AddLog(GetLogonUsername(), "DeleteUser", (*CurrentUser).String(), "")

	CurrentUser = nil
	writeJSON()
	return nil
}

// IsRegisteredUser 判断当前姓名的用户是否注册
func IsRegisteredUser(name string) bool {
	_, ok := userItems[name]
	return ok
}

// GetLogonUsername 得到当前已登录用户的姓名，如果没有登录，返回""
func GetLogonUsername() string {
	if !IsLogin() {
		return ""
	}
	return CurrentUser.Name
}

func readJSON() {
	// 解析userItems
	b1, err1 := ioutil.ReadFile(userItemsFilePath)
	if err1 == nil {
		json.Unmarshal(b1, &userItems)
	}

	// 解析CurrentUser
	b2, err2 := ioutil.ReadFile(currentUserFilePath)
	if err2 == nil {
		CurrentUser = new(Item)
		json.Unmarshal(b2, CurrentUser)
	}
}

func writeJSON() {
	// 写入userItems
	b1, err1 := json.Marshal(userItems)

	if err1 == nil {
		if _, err := os.Open(userItemsFilePath); err != nil {
			os.Create(userItemsFilePath)
		}
		ioutil.WriteFile(userItemsFilePath, b1, 0755)
	}

	// 写入CurrentUser

	if CurrentUser == nil {
		if _, err := os.Open(currentUserFilePath); err == nil {
			os.Remove(currentUserFilePath)
		}
		return
	}
	b2, err2 := json.Marshal(*CurrentUser)
	if err2 == nil {
		if _, err := os.Open(currentUserFilePath); err != nil {
			os.Create(currentUserFilePath)
		}
		ioutil.WriteFile(currentUserFilePath, b2, 0755)
	}
}

// to string
func (u Item) String() string {
	return "{Name:" + u.Name + "  Email:" + u.Email + "  Phone:" + u.PhoneNumber + "}"
}
