package gotest

import (
	"testing"

	"github.com/bilibiliChangKai/Agenda-CS/service/entity/user"
)

// 测试前，清空数据库
func TestBegin(t *testing.T) {
	user.TestClearEnv(t)
}

// user测试部分
func TestUser(t *testing.T) {
	user.TestRegisterUser(t)
	user.TestLoginLogoutUser(t)
	user.TestListUsers(t)
	user.TestDeleteUser(t)

	// 添加三个user：hza，hpz，kyj，密码都是123
	user.TestAddThreeUser(t)
}

// 测试完毕，清空数据库
func TestEnd(t *testing.T) {
	user.TestClearEnv(t)
}
