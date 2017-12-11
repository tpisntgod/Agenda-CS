package meeting

import (
	"fmt"
	"testing"
	"time"
)

var t1 = time.Now()
var t2 = t1.Add(time.Second)
var meetName = t1.Format("2006-01-02 15:04:05")

func TestCreateMeeting(t *testing.T) {
	//测试添加会议
	//时间是当前时间
	fmt.Println(t1, t2)
	var parti []string
	parti = append(parti, "aa")
	parti = append(parti, "bb")
	parti = append(parti, "cc")
	err := CreateMeeting(meetName, parti, t1, t2)
	if err != nil {
		t.Errorf("error:%s", err)
	}
}

func TestQueryMeeting(t *testing.T) {
	//测试查询会议
	err := QueryMeeting(t1, t2)
	if err != nil {
		t.Errorf("error:%s", err)
	}
}

func TestCancelMeeting(t *testing.T) {
	err := CancelMeeting(meetName)
	if err != nil {
		t.Errorf("error:%s", err)
	}
}
