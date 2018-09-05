package models

import "web-pms/util"

type Target struct {
	Id         string `json:"id" xorm:"pk varchar(50)"`
	EventId    string `json:"eventId" xorm:"varchar(50)"`
	OuterHtml  string `json:"outerHTML" xorm:"varchar(250)"`
	XPath      string `json:"XPath" xorm:"varchar(250)"`
	Selector   string `json:"selector" xorm:"varchar(250)"`
	TagSrc     string `json:"src" xorm:"varchar(250)"`
	TagName    string `json:"tagName" xorm:"varchar(50)"`
	TagId      string `json:"id" xorm:"varchar(50)"`
	TagType    string `json:"type" xorm:"varchar(50)"`
	ClassName  string `json:"className" xorm:"varchar(50)"`
	Name       string `json:"name" xorm:"varchar(50)"`
	StatusText string `json:"statusText" xorm:"varchar(50)"`
	Status     int    `json:"status" xorm:"int"`
	TimeStamp  string    `json:"timeStamp" xorm:"default '0' varchar(50)"`
}

func InsertTarget(model *Target) (int64, error) {
	var id util.UUID = util.Rand()
	model.Id = id.Hex()
	return Orm.Insert(model)
}
