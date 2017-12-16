package server

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer 新建客户端
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

// 初始化路由，分别初始化User部分和Meeting部分
func initRoutes(mx *mux.Router, formatter *render.Render) {
	initUserRoutes(mx, formatter)
	initMeetingRoute(mx, formatter)
}

// 用户部分
func initUserRoutes(mx *mux.Router, formatter *render.Render) {
	// 测试url
	mx.HandleFunc("/v1/test", test(formatter)).Methods("GET")
	// 创建用户
	mx.HandleFunc("/v1/users", createUserHandle(formatter)).Methods("POST")
	// 登录用户
	mx.HandleFunc("/v1/user/login", loginUserHandle(formatter)).Methods("POST")
	// 登出用户
	mx.HandleFunc("/v1/user/logout", logoutUserHandle(formatter)).Methods("GET")
	// 显示所有用户
	mx.HandleFunc("/v1/users", listUsersHandle(formatter)).Methods("GET")
	// 删除用户
	mx.HandleFunc("/v1/users", deleteUserHandle(formatter)).Methods("DELETE")
}

func initMeetingRoute(mx *mux.Router, formatter *render.Render) {
	// // 显示所有会议
	// mx.HandleFunc("/v1/users/{id}/all-meetings", undefinedHandler(formatter)).Methods("GET")
	// // 退出会议
	// mx.HandleFunc("/v1/users/{id}/quit-meeting/{title} ", undefinedHandler(formatter)).Methods("DELETE")
	// // 取消会议
	// mx.HandleFunc("/v1/users/{id}/cancel-meeting/{title}", undefinedHandler(formatter)).Methods("DELETE")
	// // 取消所有会议
	// mx.HandleFunc("/v1/users/cancel-all-meeting", undefinedHandler(formatter)).Methods("DELETE")
	// // 会议创建参与者
	// mx.HandleFunc("/v1/meeting/{title}/add-participators", undefinedHandler(formatter)).Methods("PUT")
	// // 会议删除参与者
	// mx.HandleFunc("/v1/meeting/{title}/delete-participators", undefinedHandler(formatter)).Methods("DELETE")
	// // 显示用户参加的所有会议
	// mx.HandleFunc("/v1/meetings", undefinedHandler(formatter)).Methods("GET")
}

func testHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id := vars["id"]
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"Hello " + id})
	}
}
