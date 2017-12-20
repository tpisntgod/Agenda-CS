package gotest

import (
	"testing"
	"time"

	"github.com/bilibiliChangKai/Agenda-CS/service/entity/meeting"
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

var t1 = time.Now()
var t2 = t1.Add(time.Second).Add(time.Second)
var startTimeConflict = t1.Add(time.Second)
var endTimeConflict = t2.Add(time.Second)
var meetName = t1.Format("2006-01-02 15:04:05")
var title = "three persons' team"

func TestCreateMeeting(t *testing.T) {
	//测试添加会议
	//时间是当前时间
	var meetingTest meeting.Meeting
	meetingTest.StartTime = t1
	meetingTest.EndTime = t2
	meetingTest.Title = "three persons' team"
	meetingTest.Host = "hza"
	meetingTest.Participator = ";kyj;"
	err := meeting.MeetingInfoService.CreateMeeting(meetingTest)
	if err != nil {
		t.Errorf("error:%s", err)
	}
}

func TestAddParticipators(t *testing.T) {
	//测试增加会议参与者
	//时间是当前时间
	var participators []string
	participators = append(participators, "hpz")
	err := meeting.MeetingInfoService.AddMeetingParticipators(title, participators)
	if err != nil {
		t.Errorf("error:%s", err)
	}
}

func TestCreateMeetingTimeConflict(t *testing.T) {
	//测试添加会议
	//时间是当前时间
	var meetingTest meeting.Meeting
	meetingTest.StartTime = startTimeConflict
	meetingTest.EndTime = endTimeConflict
	meetingTest.Title = "three persons' team time conflict"
	meetingTest.Host = "hza"
	meetingTest.Participator = ";kyj;hpz;"
	err := meeting.MeetingInfoService.CreateMeeting(meetingTest)

	if err.Error() != "user kyj time conflict,create meeting failed" {
		t.Errorf("error information can not match this :%s", err)
	}
}

func TestQueryMeeting(t *testing.T) {
	//测试增加会议参与者
	//时间是当前时间
	_, err := meeting.MeetingInfoService.QueryMeetings("hza", t1, t2)
	if err != nil {
		t.Errorf("error:%s", err)
	}
}

func TestCancelMeeting(t *testing.T) {
	//测试取消会议
	//时间是当前时间
	err := meeting.MeetingInfoService.CancelMeeting("three persons' team")
	if err != nil {
		t.Errorf("error:%s", err)
	}
}

// 测试完毕，清空数据库
func TestEnd(t *testing.T) {
	user.TestClearEnv(t)
}
