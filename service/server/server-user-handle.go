// 本层不涉及逻辑判断，逻辑判断在user.go部分

package server

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bitly/go-simplejson"
	"github.com/unrolled/render"

	"github.com/bilibiliChangKai/Agenda-CS/service/entity/user"
	"github.com/bilibiliChangKai/Agenda-CS/service/orm"
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

// error.toString
func toString(err interface{}) string {
	if err == nil {
		return ""
	}
	return fmt.Sprint(err)
}

// 标准response JSON，只包含Success和Result
func stdResj(err interface{}) resj {
	return resj{
		Information: toString(err)}
}

// 解析传过来的JSON和cookie
func praseJSON(r *http.Request) *simplejson.Json {
	// 解析json
	body, err := ioutil.ReadAll(r.Body)
	orm.CheckErr(err)
	defer r.Body.Close()

	temp, err := simplejson.NewJson(body)
	orm.CheckErr(err)
	return temp
}

func praseCookie(r *http.Request) string {
	// 解析cookie
	cookie, _ := r.Cookie("username")
	if cookie != nil {
		return cookie.Value
	}
	return ""
}

// 返回错误表单
func errResponse(w http.ResponseWriter, formatter *render.Render) {
	if err := recover(); err != nil {
		formatter.JSON(w, 500, stdResj(err))
	}
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
		defer errResponse(w, formatter)

		js := praseJSON(r)
		user.RegisterUser(
			js.Get("Name").MustString(),
			js.Get("Password").MustString(),
			js.Get("Email").MustString(),
			js.Get("Phone").MustString())

		formatter.JSON(w, http.StatusOK, stdResj(nil))
	}
}

// 登录用户
func loginUserHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer errResponse(w, formatter)

		// 使用user函数
		js := praseJSON(r)
		loginname := praseCookie(r)
		pitem := user.LoginUser(
			js.Get("Name").MustString(),
			js.Get("Password").MustString(),
			loginname)

		// 如果成功登录，设置cookie
		cookie := http.Cookie{
			Name:   "username",
			Value:  pitem.Name,
			Path:   "/",
			MaxAge: 1200}
		http.SetCookie(w, &cookie)

		resjson := stdResj(nil)
		resjson.Item = *pitem
		formatter.JSON(w, http.StatusOK, resjson)
	}
}

// 登出用户
func logoutUserHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer errResponse(w, formatter)

		loginname := praseCookie(r)
		user.LogoutUser(loginname)

		formatter.JSON(w, http.StatusOK, stdResj(nil))
	}
}

// 显示所有用户
func listUsersHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer errResponse(w, formatter)

		loginname := praseCookie(r)
		items := user.ListUsers(loginname)

		resjson := stdResj(nil)
		resjson.Users = items
		formatter.JSON(w, http.StatusOK, resjson)
	}
}

// 删除已登录用户
func deleteUserHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer errResponse(w, formatter)

		loginname := praseCookie(r)
		user.DeleteUser(loginname)

		// 如果成功删除，设置cookie
		cookie := http.Cookie{
			Name:   "username",
			Path:   "/",
			MaxAge: -1}
		http.SetCookie(w, &cookie)
		formatter.JSON(w, http.StatusOK, stdResj(nil))
	}
}

// func undefinedHandler(formatter *render.Render) http.HandlerFunc {
//
// 	return func(w http.ResponseWriter, req *http.Request) {
// 	}
// }
