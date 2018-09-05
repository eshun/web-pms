package models

import (
	"web-pms/util"
)

type Resource struct {
	Id              string `json:"id" xorm:"pk varchar(50)"`
	EventId         string `json:"eventId" xorm:"varchar(50)"`
	Name            string `json:"name" xorm:"varchar(50)"`
	Method          string `json:"method" xorm:"varchar(50)"`
	ResType         string `json:"type" xorm:"varchar(50)"`
	Duration        string `json:"duration" xorm:"default '0' varchar(50)"`
	DecodedBodySize int    `json:"decodedBodySize" xorm:"default 0 int"`
	NextHopProtocol string `json:"nextHopProtocol" xorm:"varchar(50)"`
}

func InsertResource(model *Resource) (int64, error) {
	var id util.UUID = util.Rand()
	model.Id = id.Hex()
	return Orm.Insert(model)
}

func InsertResources(list *[]Resource,eventId string) (int64, error) {
	var i int64
	var err error
	for _, v := range *list {
		v.EventId = eventId
		_, err = InsertResource(&v)
		if (err != nil) {
			return i, err
		}
		i++;
	}
	return i, err
}