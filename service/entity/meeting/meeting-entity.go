package meeting

import "time"

//Meeting 会议数据实体
type Meeting struct {
	//会议主题
	Title string `xorm:"varchar(64) pk 'title'"`
	//会议发起者
	Host string `xorm:"varchar(64) notnull 'host'"`
	//会议参与者
	Participator string `xorm:"varchar(255) 'participators'"`
	//开始时间
	StartTime time.Time `xorm:"'startTime'"`
	//结束时间
	EndTime time.Time `xorm:"'endTime'"`
}

//TableName .
func (Meeting) TableName() string {
	return "meetinginformation"
}
