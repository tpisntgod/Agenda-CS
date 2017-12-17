package meeting

import (
	"errors"
	"fmt"
	"strings"

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

//CreateMeeting 创建会议
func (*MeetingInfoAtomicService) CreateMeeting(m Meeting) error {
	meeting := new(Meeting)
	has, err := MeetingDB.Table("meetinginformation").Id(m.Title).Get(meeting)
	checkErr(err)
	//title是否已经创建
	if has {
		return errors.New("this meeting title already exists,Create Meeting failed")
	}
	participators := strings.Split(m.Participator, ";")
	fmt.Println(participators)
	//判断用户注册
	err = checkUserRegistered(participators)
	if err != nil {
		return err
	}
	//判断时间重叠
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
				return errors.New("user " + p[i] + " already joined the meeting")
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
	//check user 参加会议判断时间重叠
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

func (*MeetingInfoAtomicService) QueryMeetings() ([]Meeting, error) {

	return nil, nil
}

/*
func (*MeetingInfoAtomicService) CancelMeeting() error {

}

func (*MeetingInfoAtomicService) QuitMeeting() error {

}

func (*MeetingInfoAtomicService) ClearAllMeeting() error {

}*/
