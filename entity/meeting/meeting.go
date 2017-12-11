package meeting

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/tpisntgod/Agenda/entity/mylog"
	"github.com/tpisntgod/Agenda/entity/user"
)

type meeting_records struct {
	Title        string
	Host         string
	Participator []string
	StartTime    time.Time
	EndTime      time.Time
}

var meetings []meeting_records
var curUser string
var isFileExist = true
var writeFilePath = "src/github.com/tpisntgod/Agenda/Json/MeetingInfo.json"

//只是判断两个时间段是否overlap
func checkIfMeetingTimeOverlap(meetingStartTime, meetingEndTime, startTime, endTime time.Time) bool {
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

func checkIfTwoMeetingTimeOverlap(title string, participator []string, startTime time.Time, endTime time.Time) (int, string) {
	checkNum := 0
	var errorInfo string
	for i := 0; i < len(meetings); i++ {
		//先判断两个会议是否时间overlap，然后判断有没有同时参加这两个会议的人
		if checkIfMeetingTimeOverlap(meetings[i].StartTime, meetings[i].EndTime, startTime, endTime) {
			firstOverlap := 0
			PeopleToCheck := &meetings[i].Participator
			for j := 0; j < len(*PeopleToCheck); j++ {
				for k := 0; k < len(participator); k++ {
					if (*PeopleToCheck)[j] == participator[k] {
						if firstOverlap == 0 {
							firstOverlap = 1
							errorInfo += "以下试图添加的参与者，参加了会议" + meetings[i].Title + "，该会议的时间与想要加入的会议" + title + "存在时间冲突\n"
						}
						errorInfo += participator[k] + " "
						checkNum = 2
					}
				}
			}
		}
	}
	return checkNum, errorInfo
}

//CreateMeeting 创建会议
func CreateMeeting(title string, participator []string, startTime time.Time, endTime time.Time) error {
	//先判断两个会议是否时间overlap，然后判断有没有同时参加这两个会议的人
	for i := 0; i < len(meetings); i++ {
		if meetings[i].Title == title {
			return errors.New("会议title已经被注册")
		}
	}
	isOverlap, errorInfo := checkIfTwoMeetingTimeOverlap(title, append(participator, curUser), startTime, endTime)
	if isOverlap == 2 {
		return errors.New(errorInfo)
	}
	var meetingToAdd meeting_records
	meetingToAdd.Title = title
	meetingToAdd.Host = curUser
	for i := 0; i < len(participator); i++ {
		meetingToAdd.Participator = append(meetingToAdd.Participator, participator[i])
	}
	//将host加入participator
	meetingToAdd.Participator = append(meetingToAdd.Participator, curUser)
	meetingToAdd.StartTime = startTime
	meetingToAdd.EndTime = endTime
	meetings = append(meetings, meetingToAdd)
	mylog.AddLog(curUser, "CreateMeeting", "", meetingToAdd.String())
	WriteMeetingInfo()
	return nil
}

//AddMeetingParticipators 增加会议参与者
func AddMeetingParticipators(title string, participator []string) error {
	var isMeetingExist = 0
	for i := 0; i < len(meetings); i++ {
		if meetings[i].Title == title {
			//添加会议参与者
			isMeetingExist = 1
			//此会议不是当前用户发起
			if curUser != meetings[i].Host {
				return errors.New("此会议不是当前用户发起")
			}
			for j := 0; j < len(meetings[i].Participator); j++ {
				for k := 0; k < len(participator); k++ {
					if meetings[i].Participator[j] == participator[k] {
						return errors.New("此title的会议已经有此参与者")
					}
				}
			}
			//会议是否被注册处要修改
			//还要做时间重叠判断（允许仅有端点重叠的情况）
			isOverlap, errorInfo := checkIfTwoMeetingTimeOverlap(title, participator, meetings[i].StartTime, meetings[i].EndTime)
			if isOverlap == 2 {
				return errors.New(errorInfo)
			}
			for j := 0; j < len(participator); j++ {
				old := meetings[i]
				meetings[i].Participator = append(meetings[i].Participator, participator[j])
				mylog.AddLog(curUser, "AddParticipator", old.String(), meetings[i].String())
			}
		}
	}
	if isMeetingExist == 0 {
		return errors.New("此title的会议不存在")
	}
	WriteMeetingInfo()
	return nil
}

//DeleteMeetingParticipators 删除会议参与者
func DeleteMeetingParticipators(title string, participator []string) error {
	var isMeetingExist = 0
	var errorInfo = "以下想要删除的用户并没有参加该会议\n"
	for i := 0; i < len(meetings); i++ {
		if meetings[i].Title == title {
			isMeetingExist = 1
			//此会议不是当前用户发起
			if curUser != meetings[i].Host {
				return errors.New("此会议不是当前用户发起,当前用户没有删除权限")
			}
			var NumToDelete int
			NumToDelete = 0
			for k := 0; k < len(participator); k++ {
				isUserToDeleteExist := 0
				for j := 0; j < len(meetings[i].Participator); j++ {
					if meetings[i].Participator[j] == participator[k] {
						isUserToDeleteExist = 1
						NumToDelete++
					}
				}
				if isUserToDeleteExist == 0 {
					errorInfo += participator[k] + " "
				}
			}
			if NumToDelete < len(participator) {
				return errors.New(errorInfo)
			}
			//删除用户
			var partAfterDelete []string
			partAfterDelete = meetings[i].Participator
			for k := 0; k < len(participator); k++ {
				for j := 0; j < len(meetings[i].Participator); j++ {
					if meetings[i].Participator[j] == participator[k] {
						partAfterDelete = append(partAfterDelete[:j], partAfterDelete[j+1:]...)
					}
				}
			}
			//如果删除参与者后会议人数为0，删除该会议
			if len(partAfterDelete) == 0 {
				return CancelMeeting(title)
			}
			old := meetings[i]
			meetings[i].Participator = partAfterDelete
			mylog.AddLog(curUser, "DeleteParticipator", old.String(), meetings[i].String())
		}
	}
	if isMeetingExist == 0 {
		return errors.New("此title的会议不存在")
	}
	WriteMeetingInfo()
	return nil
}

//QueryMeeting 查询会议
func QueryMeeting(startTime time.Time, endTime time.Time) error {
	mylog.AddLog(curUser, "QueryMeeting", "", "")
	isOverlap := 0
	for i := 0; i < len(meetings); i++ {
		isCurUserInMeeting := 0
		if checkIfMeetingTimeOverlap(meetings[i].StartTime, meetings[i].EndTime, startTime, endTime) {
			for j := 0; j < len(meetings[i].Participator); j++ {
				if meetings[i].Participator[j] == curUser {
					isOverlap++
					isCurUserInMeeting = 1
					break
				}
			}
			if isCurUserInMeeting == 0 {
				continue
			}
			if isOverlap == 1 {
				fmt.Println("指定时间范围内找到的所有会议安排")
				fmt.Println("会议主题：  起始时间：  终止时间：  发起者：  参与者：")
			}
			startTimeString := meetings[i].StartTime.Format("2006-01-02 15:04:05")
			endTimeString := meetings[i].EndTime.Format("2006-01-02 15:04:05")
			fmt.Printf("%v %v %v 发起者：%v 参与者：", meetings[i].Title, startTimeString, endTimeString, meetings[i].Host)
			for j := 0; j < len(meetings[i].Participator); j++ {
				fmt.Printf("%v ", meetings[i].Participator[j])
			}
			fmt.Println()
		}
	}
	if isOverlap == 0 {
		return errors.New("此时间段并没有你的会议安排")
	}
	return nil
}

//QueryMeetingWebVersion 查询会议
//这个不是第一次作业的部分，是开发web服务程序作业的函数
//因为之前没有考虑到web服务的Agenda，所以这里补充了一个函数
func QueryMeetingWebVersion(startTime time.Time, endTime time.Time) (string, error) {
	var queryResult string
	mylog.AddLog(curUser, "QueryMeeting", "", "")
	isOverlap := 0
	for i := 0; i < len(meetings); i++ {
		isCurUserInMeeting := 0
		if checkIfMeetingTimeOverlap(meetings[i].StartTime, meetings[i].EndTime, startTime, endTime) {
			for j := 0; j < len(meetings[i].Participator); j++ {
				if meetings[i].Participator[j] == curUser {
					isOverlap++
					isCurUserInMeeting = 1
					break
				}
			}
			if isCurUserInMeeting == 0 {
				continue
			}
			if isOverlap == 1 {
				queryResult += "指定时间范围内找到的所有会议安排\n"
				queryResult += "会议主题：  起始时间：  终止时间：  发起者：  参与者：\n"
			}
			startTimeString := meetings[i].StartTime.Format("2006-01-02 15:04:05")
			endTimeString := meetings[i].EndTime.Format("2006-01-02 15:04:05")
			queryResult += fmt.Sprintf("%v %v %v 发起者：%v 参与者：", meetings[i].Title, startTimeString, endTimeString, meetings[i].Host)
			for j := 0; j < len(meetings[i].Participator); j++ {
				queryResult += fmt.Sprintf("%v ", meetings[i].Participator[j])
			}
			queryResult += "\n"
		}
	}
	if isOverlap == 0 {
		return queryResult, errors.New("此时间段并没有你的会议安排")
	}
	return queryResult, nil
}

//CancelMeeting 取消会议
func CancelMeeting(title string) error {
	isMeetingExist := 0
	for i := 0; i < len(meetings); i++ {
		if meetings[i].Title == title {
			isMeetingExist = 1
			if meetings[i].Host == curUser {
				mylog.AddLog(curUser, "CancelMeeting", meetings[i].String(), "")
				meetings = append(meetings[:i], meetings[i+1:]...)
				break
			} else {
				fmt.Println("not host")
				return errors.New("此会议的发起者并不是当前用户,当前用户没有取消会议权限")
			}
		}
	}
	if isMeetingExist == 0 {
		return errors.New("此主题的会议不存在")
	}
	WriteMeetingInfo()
	return nil
}

//QuitMeeting 退出会议
func QuitMeeting(title string) error {
	isMeetingExist := 0
	isAttendMeeting := 0
	for i := 0; i < len(meetings); i++ {
		if meetings[i].Title == title {
			isMeetingExist = 1
			for j := 0; j < len(meetings[i].Participator); j++ {
				if meetings[i].Participator[j] == curUser {
					isAttendMeeting = 1
					meetings[i].Participator = append(meetings[i].Participator[:j], meetings[i].Participator[j+1:]...)
					if len(meetings[i].Participator) == 0 {
						return CancelMeeting(title)
					}
					break
				}
			}
			mylog.AddLog(curUser, "QuitMeeting", meetings[i].String(), "")
			break
		}
	}
	if isMeetingExist == 0 {
		return errors.New("此主题的会议不存在")
	}
	if isAttendMeeting == 0 {
		return errors.New("你并没有参加此会议")
	}
	WriteMeetingInfo()
	return nil
}

//ClearAllMeeting 清空会议
func ClearAllMeeting() {
	for i := 0; i < len(meetings); i++ {
		if meetings[i].Host == curUser {
			meetings = append(meetings[:i], meetings[i+1:]...)
			i--
		}
	}
	mylog.AddLog(curUser, "ClearMeeting", "", "")
	WriteMeetingInfo()
}

//WriteMeetingInfo 将会议信息以JSON格式写入文件
func WriteMeetingInfo() {
	b, err := json.Marshal(meetings)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(writeFilePath, b, 0644)
	if err != nil {
		panic(err)
	}
}

//SetCurrentUser 设置当前用户
func SetCurrentUser(currentUser string) {
	curUser = currentUser
}

//CheckStarttimelessthanEndtime 判断输入的startTime是否小于endTime
func CheckStarttimelessthanEndtime(startTime time.Time, endTime time.Time) bool {
	if startTime.After(endTime) {
		return false
	}
	return true
}

func init() {
	//根据Current.txt知道当前登录用户
	curUser = user.GetLogonUsername()
	//writeFilePath = filepath.Join(os.Getenv("GOPATH"), writeFilePath)
	writeFilePath = filepath.Join(*mylog.GetGOPATH(), writeFilePath)
	_, err2 := os.Stat(writeFilePath)
	if err2 == nil {
		data, err := ioutil.ReadFile(writeFilePath)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(data, &meetings)
		if err != nil {
			panic(err)
		}
	} else {
		if os.IsNotExist(err2) {
			isFileExist = false
		} else {
			fmt.Println(errors.New("保存会议信息的文件打开失败"))
		}
	}
}

// to string
func (m meeting_records) String() string {
	var outStr = "{Title:" + m.Title + "  Host:" + m.Host + "  Participor:"
	for _, s := range m.Participator {
		outStr += s + " "
	}
	outStr += " Begin Time:" + m.StartTime.Format("2006-01-02 15:04:05") + "  End Time:" + m.EndTime.Format("2006-01-02 15:04:05") + "}"
	return outStr
}
