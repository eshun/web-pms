package models

import "web-pms/util"

type Req struct {
	Id      string `json:"id" xorm:"pk varchar(50)"`
	EventId string `json:"eventId" xorm:"varchar(50)"`
	Url     string `json:"url" xorm:"varchar(250)"`
	Method  string `json:"method" xorm:"varchar(50)"`
}

func InsertReq(model *Req) (int64, error) {
	var id util.UUID = util.Rand()
	model.Id = id.Hex()
	return Orm.Insert(model)
}
