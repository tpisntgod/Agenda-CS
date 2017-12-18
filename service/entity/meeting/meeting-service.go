package meeting

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/bilibiliChangKai/Agenda-CS/service/entity/user"
)

//MeetingInfoAtomicService .
type MeetingInfoAtomicService struct{}

//MeetingInfoService .
var MeetingInfoService = MeetingInfoAtomicService{}

//meetingjson 创建会议 存放json解析后的数据
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

func checkUserRegistered(p []string) error {
	fmt.Println("checkUserRegistered")
	for i := 0; i < len(p); i++ {
		user := new(user.Item)
		has, err := MeetingDB.Table("item").Id(p[i]).Get(user)
		if has == false {
			return errors.New("user " + p[i] + " didn't register")
		}
		if err != nil {
			return err
		}
	}
	return nil
}

//只是判断两个时间段是否overlap
func checkIfMeetingTimeOverlap(meetingStartTime, meetingEndTime, startTime, endTime time.Time) bool {
	fmt.Println("checkIfMeetingTimeOverlap")
	fmt.Println(meetingStartTime, meetingEndTime, startTime, endTime)
	if (meetingStartTime.Before(startTime) || meetingStartTime.Equal(startTime)) &&
		meetingEndTime.After(startTime) && (meetingEndTime.Before(endTime) || meetingEndTime.Equal(endTime)) {
		return true
	}
	if (meetingStartTime.Before(startTime) || meetingStartTime.Equal(startTime)) &&
		(meetingEndTime.After(endTime) || meetingEndTime.Equal(endTime)) {
		return true
	}
	if (meetingStartTime.After(startTime) || meetingStartTime.Equal(startTime)) &&
		meetingStartTime.Before(endTime) && (meetingEndTime.After(endTime) || meetingEndTime.Equal(endTime)) {
		return true
	}
	if (meetingStartTime.After(startTime) || meetingStartTime.Equal(startTime)) &&
		(meetingEndTime.Before(endTime) || meetingEndTime.Equal(endTime)) {
		return true
	}
	return false
}

//CreateMeeting 创建会议
func (*MeetingInfoAtomicService) CreateMeeting(m Meeting) error {
	fmt.Println("CreateMeeting")
	meeting := new(Meeting)
	has, err := MeetingDB.Table("meetinginformation").Id(m.Title).Get(meeting)
	checkErr(err)
	//title是否已经创建
	if has {
		return errors.New("this meeting title already exists,create meeting failed")
	}
	participators := strings.Split(m.Participator, ";")
	fmt.Println(participators)
	//判断用户注册
	err = checkUserRegistered(participators)
	if err != nil {
		return err
	}
	//判断参加会议的用户时间是否重叠
	//作为会议参与者相关的会议信息
	for i := 0; i < len(participators); i++ {
		sql := "select * from meetinginformation where participators like '%" + participators[i] + "%'"
		result, err := GetMeetingInTimeInterval(1, participators[i], sql, m.StartTime, m.EndTime)
		if result != "" {
			return errors.New("user " + participators[i] + " time conflict,create meeting failed")
		}
		if err != nil {
			return err
		}
	}

	_, err = MeetingDB.Table("meetinginformation").Insert(m)
	checkErr(err)
	if err == nil {
		fmt.Println("create success")
		return nil
	}
	fmt.Println("create error")
	return errors.New("meeting insert failed")
}

//增加参加者时，判断用户是否已经加入过该会议
func checkUserAlreadyJoin(alreadyjoin []string, p []string) error {
	for i := 0; i < len(p); i++ {
		for j := 0; j < len(alreadyjoin); j++ {
			if p[i] == alreadyjoin[j] {
				return errors.New("user " + p[i] + " already joined the meeting,add participators failed")
			}
		}
	}
	return nil
}

//AddMeetingParticipators 增加会议参与者
func (*MeetingInfoAtomicService) AddMeetingParticipators(title string, p []string) error {
	//meeting := Meeting{Title: title}
	meeting := new(Meeting)
	has, err := MeetingDB.Id(title).Get(meeting)
	checkErr(err)
	if has == false {
		return errors.New("this meeting doesn't exist,add participator failed")
	}
	alreadyJoin := strings.Split(meeting.Participator, ";")
	err = checkUserRegistered(p)
	if err != nil {
		return err
	}
	err = checkUserAlreadyJoin(alreadyJoin, p)
	if err != nil {
		return err
	}
	//判断参加会议的用户时间是否重叠
	//作为会议参与者相关的会议信息
	for i := 0; i < len(p); i++ {
		sql := "select * from meetinginformation where participators like '%" + p[i] + "%'"
		result, err := GetMeetingInTimeInterval(1, p[i], sql, meeting.StartTime, meeting.EndTime)
		if result != "" {
			return errors.New("user " + p[i] + " time conflict,add participator failed")
		}
		if err != nil {
			return err
		}
	}

	var add string
	for i := 0; i < len(p); i++ {
		add += ";" + p[i]
	}
	meeting.Participator += add
	fmt.Println("add participators")
	fmt.Println(meeting)

	_, err = MeetingDB.Id(title).Update(meeting)
	if err != nil {
		return err
	}
	return nil
}

func (*MeetingInfoAtomicService) DeleteMeetingParticipators() error {
	return nil
}

func CheckUserInMeeting(queryType int, host string, participators string, user string) int {
	//判断该用户是否真的在会议中
	var usercheck int
	//queryType是2表明判断用户是不是host，1表明判断用户是不是participator
	if queryType == 2 {
		if host == user {
			return 1
		}
		return 0
	}
	if strings.Contains(participators, ";"+user+";") {
		usercheck++
	}
	if strings.Contains(participators, user+";") {
		usercheck++
	}
	if strings.Contains(participators, ";"+user) {
		usercheck++
	}
	if participators == user {
		usercheck++
	}
	return usercheck
}

var checkDup map[string]int

//GetMeetingInTimeInterval 判断用户加入的会议是否与此时间段冲突
//返回值第一个是""的话，表示没有会议时间重叠
func GetMeetingInTimeInterval(queryType int, user string, sql string, starttime time.Time, endTime time.Time) (string, error) {
	var userInMeeting int
	results, err := MeetingDB.Query(sql)
	if err != nil {
		return "", err
	}
	if queryType == 1 {
		checkDup = make(map[string]int)
	}
	var meetingInfo string
	for _, k := range results {
		title := string(k["title"])
		//判断会议是否重复加入
		if checkDup[title] == 1 {
			continue
		}
		if CheckUserInMeeting(queryType, string(k["host"]), string(k["participators"]), user) == 0 {
			continue
		}
		checkDup[title] = 1
		stime := string(k["startTime"])
		stime = strings.Replace(stime, "T", " ", 1)
		stime = strings.Replace(stime, "Z", "", 1)
		etime := string(k["endTime"])
		etime = strings.Replace(etime, "T", " ", 1)
		etime = strings.Replace(etime, "Z", "", 1)
		startTimeFormat, _ := time.Parse("2006-01-02 15:04:05", stime)
		endTimeFormat, _ := time.Parse("2006-01-02 15:04:05", etime)
		//fmt.Println("s22t", stime, etime)
		if checkIfMeetingTimeOverlap(startTimeFormat, endTimeFormat, starttime, endTime) {
			userInMeeting++
			meetingInfo += title + "    " + stime + "  " + etime + "  " +
				string(k["host"]) + "    " + string(k["participators"]) + "\n"
		}
	}
	if userInMeeting == 0 {
		return "", nil
	}
	return meetingInfo, nil
}

func (*MeetingInfoAtomicService) QueryMeetings(user string, starttime time.Time, endTime time.Time) (string, error) {
	//作为会议参与者相关的会议信息
	meetingInfoTitle := "指定时间范围内找到的所有会议安排\n会议主题： 起始时间：      终止时间：      发起者：  参与者：      \n"
	var meetingInfo string
	//作为会议参与者相关的会议信息
	sql := "select * from meetinginformation where participators like '%" + user + "%'"
	result, err := GetMeetingInTimeInterval(1, user, sql, starttime, endTime)
	meetingInfo += result
	if err != nil {
		return "", err
	}
	//作为会议主持者相关的会议信息
	sql2 := "select * from meetinginformation where host like '%" + user + "%'"
	result, err = GetMeetingInTimeInterval(2, user, sql2, starttime, endTime)
	meetingInfo += result
	if err != nil {
		return "", err
	}
	//用户什么会议都没参加/主持
	if meetingInfo == "" {
		return "", errors.New("user" + user + "participate no meetings")
	}
	meetingInfo = meetingInfoTitle + meetingInfo
	fmt.Println("meetingInfo:")
	fmt.Println(meetingInfo)
	return meetingInfo, nil
}

/*
func (*MeetingInfoAtomicService) CancelMeeting() error {

}

func (*MeetingInfoAtomicService) QuitMeeting() error {

}

func (*MeetingInfoAtomicService) ClearAllMeeting() error {

}*/
