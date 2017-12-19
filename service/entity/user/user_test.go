package user

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"testing"
	"time"

	"github.com/bilibiliChangKai/Agenda-CS/service/orm"
)

//hash the password
func hashFunc(hashString string) string {
	h := md5.New()
	h.Write([]byte(hashString))
	cipheStr := h.Sum(nil)
	return hex.EncodeToString(cipheStr)
}

func TestBuildEnv(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()

	_, err := orm.Mydb.Exec("DELETE FROM item")
	checkErr(err)

	_, err = orm.Mydb.Unscoped().Exec("DELETE FROM item")
	checkErr(err)
}
func TestRegisterUser(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()

	// 测试注册用户
	// 一共注册20个用户
	for i := 0; i < 20; i++ {
		RegisterUser(
			"Name"+string(i),
			hashFunc("Password"+string(i)),
			"12345678@qq.com",
			"12345678910")
	}
}
func TestLoginLogoutUser(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()

	// 测试登录登出
	// 依次登录登出20个用户
	for i := 0; i < 20; i++ {
		LoginUser("Name"+string(i), hashFunc("Password"+string(i)), "")
		LogoutUser("Name" + string(i))
	}
}
func TestListUsers(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()

	// 测试列出用户
	// 依次测试列出的用户
	items := ListUsers("Name1")
	if len(items) != 20 {
		t.Error("Error : User number is not equal 20!")
	}
	for i := 0; i < 20; i++ {
		if items[i].Email != "12345678@qq.com" {
			t.Errorf("Error : User phone is wrong!")
		}
		if items[i].PhoneNumber != "12345678910" {
			t.Errorf("Error : User phone is wrong!")
		}
	}
}
func TestDeleteUser(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Error("Error : User isn't deleted!")
		}
	}()

	// 测试删除
	// 删除全部用户
	for i := 0; i < 20; i++ {
		DeleteUser("Name" + string(i))
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := r.Intn(100)
	LoginUser("Name"+string(i), "Password"+string(i), "")
}
func TestAddThreeUser(t *testing.T) {
	RegisterUser("hza", "123", "123456@qq.com", "123456789")
	RegisterUser("hpz", "123", "123789@qq.com", "123456710")
	RegisterUser("kyj", "123", "654321@qq.com", "123456712")
}
