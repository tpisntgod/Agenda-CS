package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/bilibiliChangKai/Agenda-CS/service/entity/meeting"
	"github.com/unrolled/render"
)

type resjson struct {
	Meetings    []meeting.Meeting
	Information string
}

//meetingjson 用于json解析的会议数据实体
type meetingjson struct {
	//会议主题
	Title string
	//会议发起者
	Host string
	//会议参与者
	Participator []string
	//开始时间
	StartTime string
	//结束时间
	EndTime string
}

var meetingService = meeting.MeetingInfoService

// 将参与者名字的类型[]string转成string方便数据库存储
func getParticipatorsName(p []string) string {
	s := ""
	for i := 0; i < len(p)-1; i++ {
		s = s + p[i] + ";"
	}
	s = s + p[len(p)-1]
	return s
}

// 返回cookie中携带的Name字段
func getCurrentUserNameMeeting(r *http.Request) string {
	cookie, _ := r.Cookie("username")
	if cookie != nil {
		return cookie.Value
	} else {
		fmt.Println("cookie nil")
	}
	return "unknown"
}

//getResponseJson 构造http response的json
func getResponseJson(info string) resjson {
	return resjson{
		Information: info}
}

//创建会议 /v1/meetings
func createMeetingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		/*
			bodyString := string(body)
			fmt.Println(bodyString)*/

		var meetingj meetingjson
		//var meeting meeting.Meeting
		if err := json.Unmarshal(body, &meetingj); err == nil {
			fmt.Println(meetingj)

			starttime, _ := time.Parse("2006-01-02 15:04:05", meetingj.StartTime)
			endtime, _ := time.Parse("2006-01-02 15:04:05", meetingj.EndTime)
			meeting := meeting.Meeting{Title: meetingj.Title, Host: getCurrentUserNameMeeting(r),
				Participator: getParticipatorsName(meetingj.Participator), StartTime: starttime, EndTime: endtime}

			fmt.Println("meetingj.Participator")
			for i := 0; i < len(meetingj.Participator); i++ {
				fmt.Println(meetingj.Participator[i])
			}

			fmt.Println("meeting")
			fmt.Println(meeting)

			err := meetingService.CreateMeeting(meeting)
			var info string
			if err != nil {
				info = toString(err)
			} else {
				info = "create meeting succeed"
			}
			formatter.JSON(w, http.StatusOK, getResponseJson(info))
			/*
				ret, _ := json.Marshal(meetingj)
				fmt.Fprint(w, "\n")
				fmt.Fprint(w, string(ret))*/
		} else {
			fmt.Println(err)
		}

		return
		/*
			r.ParseForm()
			participators := getParticipatorsName(r.Form["participators"])
			starttime, _ := time.Parse("2006-01-02 15:04:05", r.Form["stime"][0])
			endtime, _ := time.Parse("2006-01-02 15:04:05", r.Form["etime"][0])
			m := meeting.Meeting{Title: r.Form["title"][0], Host: getCurrentUserName(r),
				Participator: participators, StartTime: starttime, EndTime: endtime}
			err := meetingService.CreateMeeting(m)
			info := toString(err)
			formatter.JSON(w, http.StatusOK, getResponseJson(info))*/
	}
}

//增加会议参与者 /v1/meeting/{title}/adding-participators
func addParticipatorsHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		if len(r.Form["title"]) == 0 {
			fmt.Println("parse error")
		} else {
			fmt.Println("title:", r.Form["title"][0])
		}
	}
}

//删除会议参与者 /v1/meeting/{title}/deleting-participators
func deleteParticipatorsHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

//查询会议 /v1/users/query-meeting{?starttime,endtime}
func queryMeetingsHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

//取消会议 /v1/users/cancel-a-meeting/{title}
func cancelMeetingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

//退出会议 /v1/users/quit-meeting/{title}
func quitMeetingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

//清空会议 /v1/users/cancel-all-meeting
func clearAllMeetingsHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("clear all")
	}
}
