package user

// Item 用户信息
type Item struct {
	// 用户名字，是唯一主键
	Name string `xorm:"pk" json:",omitempty"`
	// hash过的密码
	HashPassword string `json:",-"`
	// 注册用的邮箱
	Email string `json:",omitempty"`
	// 注册用的电话号码
	PhoneNumber string `json:"Phone,omitempty"`
}

// newItem 新建一个UserItem，并返回指针
func newItem(name string, password string,
	email string, phoneNumber string) *Item {
	newUserItem := new(Item)
	newUserItem.Name = name
	newUserItem.HashPassword = password
	newUserItem.Email = email
	newUserItem.PhoneNumber = phoneNumber
	return newUserItem
}

// // 用于密码hash的函数
// func hashFunc(hashString string) string {
// 	// 进行md5加密
// 	h := md5.New()
// 	h.Write([]byte(hashString))
// 	cipherStr := h.Sum(nil)
// 	return hex.EncodeToString(cipherStr)
// }
