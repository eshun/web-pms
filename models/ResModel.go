package models

import "web-pms/util"

type Res struct {
	Id          string `json:"id" xorm:"pk varchar(50)"`
	EventId     string `json:"eventId" xorm:"varchar(50)"`
	Response    string `json:"response" xorm:"varchar(250)"`
	ElapsedTime int    `json:"elapsedTime" xorm:"int"`
	StatusText  string `json:"statusText" xorm:"varchar(50)"`
	Status      int    `json:"status" xorm:"int"`
}

func InsertRes(model *Res) (int64, error) {
	var id util.UUID = util.Rand()
	model.Id = id.Hex()
	return Orm.Insert(model)
}
