package meeting

import (
	"testing"
	"time"
)

var t1 = time.Now()
var t2 = t1.Add(time.Second).Add(time.Second)
var startTimeConflict = t1.Add(time.Second)
var endTimeConflict = t2.Add(time.Second)
var meetName = t1.Format("2006-01-02 15:04:05")
var title = "three persons' team"

func TestCreateMeeting(t *testing.T) {
	//测试添加会议
	//时间是当前时间
	var meeting Meeting
	meeting.StartTime = t1
	meeting.EndTime = t2
	meeting.Title = "three persons' team"
	meeting.Host = "hza"
	meeting.Participator = ";kyj;"
	err := MeetingInfoService.CreateMeeting(meeting)
	if err != nil {
		t.Errorf("error:%s", err)
	}
}

func TestAddParticipators(t *testing.T) {
	//测试增加会议参与者
	//时间是当前时间
	var participators []string
	participators = append(participators, "hpz")
	err := MeetingInfoService.AddMeetingParticipators(title, participators)
	if err != nil {
		t.Errorf("error:%s", err)
	}
}

func TestCreateMeetingTimeConflict(t *testing.T) {
	//测试添加会议
	//时间是当前时间
	var meeting Meeting
	meeting.StartTime = startTimeConflict
	meeting.EndTime = endTimeConflict
	meeting.Title = "three persons' team time conflict"
	meeting.Host = "hza"
	meeting.Participator = ";kyj;hpz;"
	err := MeetingInfoService.CreateMeeting(meeting)

	if err.Error() != "user kyj time conflict,create meeting failed" {
		t.Errorf("error information can not match this :%s", err)
	}
}

func TestQueryMeeting(t *testing.T) {
	//测试增加会议参与者
	//时间是当前时间
	_, err := MeetingInfoService.QueryMeetings("hza", t1, t2)
	if err != nil {
		t.Errorf("error:%s", err)
	}
}

func TestCancelMeeting(t *testing.T) {
	//测试取消会议
	//时间是当前时间
	err := MeetingInfoService.CancelMeeting("three persons' team")
	if err != nil {
		t.Errorf("error:%s", err)
	}
}
