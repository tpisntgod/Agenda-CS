package meeting

import (
	"errors"
	"fmt"
	"strings"
<<<<<<< HEAD

	"github.com/bilibiliChangKai/Agenda-CS/service/entity/user"
=======
>>>>>>> ccf4b8ea97c1f569b12ae370caa8b1a3855d292a
)

//MeetingInfoAtomicService .
type MeetingInfoAtomicService struct{}

//MeetingInfoService .
var MeetingInfoService = MeetingInfoAtomicService{}

<<<<<<< HEAD
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

=======
>>>>>>> ccf4b8ea97c1f569b12ae370caa8b1a3855d292a
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
<<<<<<< HEAD
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

=======
	fmt.Println("create err")
	return errors.New("meeting title already exists")
}

>>>>>>> ccf4b8ea97c1f569b12ae370caa8b1a3855d292a
//AddMeetingParticipators 增加会议参与者
func (*MeetingInfoAtomicService) AddMeetingParticipators(title string, p []string) error {
	//meeting := Meeting{Title: title}
	meeting := new(Meeting)
	has, err := MeetingDB.Id(title).Get(meeting)
<<<<<<< HEAD
=======
	fmt.Println("get meeting")
	fmt.Println(meeting)
	alreadyJoin := strings.Split(meeting.Participator, ";")
	fmt.Println(alreadyJoin)
>>>>>>> ccf4b8ea97c1f569b12ae370caa8b1a3855d292a
	checkErr(err)
	if has == false {
		return errors.New("this meeting doesn't exist,add participator failed")
	}
<<<<<<< HEAD
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
=======
>>>>>>> ccf4b8ea97c1f569b12ae370caa8b1a3855d292a
	var add string
	for i := 0; i < len(p); i++ {
		add += ";" + p[i]
	}
	meeting.Participator += add
	fmt.Println("add participators")
	fmt.Println(meeting)
<<<<<<< HEAD

	_, err = MeetingDB.Id(title).Update(meeting)
	if err != nil {
		return err
	}
=======
	/*
		_, err = MeetingDB.Id(title).Update(meeting)
		if err != nil {
			return err
		}*/
>>>>>>> ccf4b8ea97c1f569b12ae370caa8b1a3855d292a
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
