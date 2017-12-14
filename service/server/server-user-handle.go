package server

import (
	"net/http"

	"github.com/unrolled/render"

	"github.com/tpisntgod/Agenda/service/entity/user"
)

// 用于返回的模板Json
type resj struct {
	// 包含userItem属性
	user.Item
	// 返回user查询列表
	Items []user.Item
	// 结果是否成功
	Success bool
	// 表示结果
	Result string
}

// 通过user中的函数，更新CurrentUser
func updateCurrentUser(r *http.Request) {
	cookie, err := r.Cookie("Name")
	user.SetCurrentUser(cookie.Value, err)
}

// error.toString
func toString(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

// 标准response JSON，只包含Success和Result
func stdResj(succ bool, res string) resj {
	return resj{
		Success: succ,
		Result:  res}
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

// 显示CurrentUser的id和name
func test(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resjson := resj{}
		//resjson.Items = append(resjson.Items, user.Item{"1", "2", "3", "4"})
		resjson.Name = "123"
		resjson.Success = true
		resjson.Result = "trytry"
		formatter.JSON(w, http.StatusOK, resjson)
	}
}

// 创建一个新的用户
func createUserHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		err := user.RegisterUser(r.FormValue("Name"), r.FormValue("Passname"), r.FormValue("Email"), r.FormValue("Phone"))
		succ := (bool)(err == nil)
		res := toString(err)

		formatter.JSON(w, http.StatusOK, struct {
			Success bool
			Result  string
		}{
			succ,
			res})
	}
}

// 登录用户
func loginUserHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 使用user函数
		updateCurrentUser(r)
		r.ParseForm()
		err := user.LoginUser(r.FormValue("Name"), r.FormValue("Password"))
		succ := (bool)(err == nil)
		res := toString(err)

		// 返回报文
		if succ {
			resjson := stdResj(succ, res)
			resjson.Item = *user.CurrentUser
			formatter.JSON(w, http.StatusOK, resjson)
		} else {
			formatter.JSON(w, http.StatusOK, stdResj(succ, res))
		}
	}
}

// 登出用户
func logoutUserHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 使用user函数
		updateCurrentUser(r)
		r.ParseForm()
		err := user.LogoutUser()
		succ := (bool)(err == nil)
		res := toString(err)
		formatter.JSON(w, http.StatusOK, stdResj(succ, res))
	}
}

func undefinedHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
	}
}
