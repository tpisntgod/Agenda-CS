// 本层不涉及逻辑判断，逻辑判断在user.go部分

package server

import (
	"fmt"
	"net/http"

	"github.com/unrolled/render"

	"github.com/bilibiliChangKai/Agenda-CS/service/entity/user"
)

// 用于返回的模板Json
type resj struct {
	// 包含userItem属性
	user.Item
	// 返回user查询列表
	Users []user.Item `json:",omitempty"`
	// 表示结果
	Information string
}

// 返回cookie中携带的Name字段
func getCurrentUserName(r *http.Request) string {
	cookie, err := r.Cookie("Name")
	if cookie != nil {
		return cookie.Value
	} else {
		return ""
	}
}

// error.toString
func toString(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

// 标准response JSON，只包含Success和Result
func stdResj(inf string) resj {
	return resj{
		Information: inf}
}

func initMydb(args []string) {
	// if len(args) != 5 && len(args) != 1 {
	// 	fmt.Fprintln(os.Stderr, "Please input the database information!")
	// 	fmt.Fprintln(os.Stderr, "\t./app username password port databasename")
	// 	fmt.Fprintln(os.Stderr, "Or use: \n\t./app\nwe will use (root) (root) (2048) (test)")
	// 	os.Exit(1)
	// }
	//
	// // 声明四个变量
	// name := "root"
	// password := "root"
	// port := "2048"
	// dname := "test"
	//
	// if len(args) != 1 {
	// 	name = args[1]
	// 	password = args[2]
	// 	port = args[3]
	// 	dname = args[4]
	// }
	//
	// // 创建数据库
	// entities.InitMydb(name, password, port, dname)
}

// test
func test(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resjson := resj{}
		resjson.Users = append(resjson.Users, user.Item{"1", "2", "3", "4"})
		resjson.Name = "123"
		resjson.Email = "321"
		resjson.PhoneNumber = "456"
		resjson.Information = "trytry"
		formatter.JSON(w, http.StatusOK, resjson)
	}
}

// 创建一个新的用户
func createUserHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := user.RegisterUser(r.FormValue("Name"), r.FormValue("Password"), r.
			FormValue("Email"), r.FormValue("Phone"))
		//succ := (bool)(err == nil)
		res := toString(err)
		formatter.JSON(w, http.StatusOK, stdResj(res))
	}
}

// 登录用户
func loginUserHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 使用user函数
		loginname := getCurrentUserName(r)
		r.ParseForm()
		pitem, err := user.LoginUser(r.FormValue("Name"), r.FormValue("Password"), loginname)
		succ := (bool)(err == nil)
		res := toString(err)

		// 返回报文
		if succ {
			// 如果成功登录，设置cookie
			cookie := http.Cookie{
				Name:   "Name",
				Value:  pitem.Name,
				Path:   "/",
				MaxAge: 1200}
			http.SetCookie(w, &cookie)

			resjson := stdResj(res)
			resjson.Item = *pitem
			formatter.JSON(w, http.StatusOK, resjson)
		} else {
			formatter.JSON(w, http.StatusOK, stdResj(res))
		}
	}
}

// 登出用户
func logoutUserHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		loginname := getCurrentUserName(r)
		err := user.LogoutUser(loginname)
		//succ := (bool)(err == nil)
		res := toString(err)
		formatter.JSON(w, http.StatusOK, stdResj(res))
	}
}

// 显示所有用户
func listUsersHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		loginname := getCurrentUserName(r)
		fmt.Println(r.Cookies())
		items, err := user.ListUsers(loginname)
		//succ := (bool)(err == nil)
		res := toString(err)
		if items == nil {
			formatter.JSON(w, http.StatusOK, stdResj(res))
		} else {
			resjson := stdResj(res)
			resjson.Users = items
			formatter.JSON(w, http.StatusOK, resjson)
		}
	}
}

// 删除已登录用户
func deleteUserHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		loginname := getCurrentUserName(r)
		err := user.DeleteUser(loginname)
		//succ := (bool)(err == nil)
		res := toString(err)

		if err == nil {
			// 如果成功删除，设置cookie
			cookie := http.Cookie{
				Name:   "Name",
				Path:   "/",
				MaxAge: -1}
			http.SetCookie(w, &cookie)
		}
		formatter.JSON(w, http.StatusOK, stdResj(res))
	}
}

func undefinedHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
	}
}
