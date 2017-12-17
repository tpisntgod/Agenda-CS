package meeting

import (
	"errors"
	"fmt"
	"strings"
)

//MeetingInfoAtomicService .
type MeetingInfoAtomicService struct{}

//MeetingInfoService .
var MeetingInfoService = MeetingInfoAtomicService{}

//CreateMeeting 创建会议
func (*MeetingInfoAtomicService) CreateMeeting(m Meeting) error {
	_, err := MeetingDB.Table("meetinginformation").Insert(m)
	checkErr(err)
	if err == nil {
		fmt.Println("create success")
		return nil
	}
	fmt.Println("create err")
	return errors.New("meeting title already exists")
}

//AddMeetingParticipators 增加会议参与者
func (*MeetingInfoAtomicService) AddMeetingParticipators(title string, p []string) error {
	//meeting := Meeting{Title: title}
	meeting := new(Meeting)
	has, err := MeetingDB.Id(title).Get(meeting)
	fmt.Println("get meeting")
	fmt.Println(meeting)
	alreadyJoin := strings.Split(meeting.Participator, ";")
	fmt.Println(alreadyJoin)
	checkErr(err)
	if has == false {
		return errors.New("this meeting doesn't exist,add participator failed")
	}
	var add string
	for i := 0; i < len(p); i++ {
		add += ";" + p[i]
	}
	meeting.Participator += add
	fmt.Println("add participators")
	fmt.Println(meeting)
	/*
		_, err = MeetingDB.Id(title).Update(meeting)
		if err != nil {
			return err
		}*/
	return nil
}

func (*MeetingInfoAtomicService) DeleteMeetingParticipators() error {
	return nil
}

/*
func (*MeetingInfoAtomicService) QueryMeetings() ([]Meeting, error) string {

}

func (*MeetingInfoAtomicService) CancelMeeting() error {

}

func (*MeetingInfoAtomicService) QuitMeeting() error {

}

func (*MeetingInfoAtomicService) ClearAllMeeting() error {

}*/
