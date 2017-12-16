package meeting

import (
	"fmt"
)

//MeetingInfoAtomicService .
type MeetingInfoAtomicService struct{}

//MeetingInfoService .
var MeetingInfoService = MeetingInfoAtomicService{}

func (*MeetingInfoAtomicService) CreateMeeting(m Meeting) error {
	_, err := MeetingDB.Insert(m)
	checkErr(err)
	if err == nil {
		fmt.Println("create success")
	} else {
		fmt.Println("create err")
	}
	return err
}

func (*MeetingInfoAtomicService) AddMeetingParticipators() error {
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
