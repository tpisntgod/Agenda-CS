package user

import (
	"crypto/md5"
	"encoding/hex"
)

// UserItem 用户信息
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

// NewUserItem 新建一个Item，并返回指针
func NewUserItem(name string, password string,
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
