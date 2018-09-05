package models

import "web-pms/util"

type PageTimes struct {
	Id              string `json:"id" xorm:"pk varchar(50)"`
	EventId         string `json:"eventId" xorm:"varchar(50)"`
	DnsTime         int    `json:"dnsTime" xorm:"default 0 int"`
	TcpTime         int    `json:"tcpTime" xorm:"default 0 int"`
	WhiteTime       int    `json:"whiteTime" xorm:"default 0 int"`
	DomTime         int    `json:"domTime" xorm:"default 0 int"`
	LoadTime        int    `json:"loadTime" xorm:"default 0 int"`
	ReadyTime       int    `json:"readyTime" xorm:"default 0 int"`
	RedirectTime    int    `json:"redirectTime" xorm:"default 0 int"`
	UnloadTime      int    `json:"unloadTime" xorm:"default 0 int"`
	RequestTime     int    `json:"requestTime" xorm:"default 0 int"`
	AnalysisDomTime int    `json:"analysisDomTime" xorm:"default 0 int"`
	ResourceTime    int    `json:"resourceTime" xorm:"default 0 int"`
}

func InsertPageTimes(model *PageTimes) (int64, error) {
	var id util.UUID = util.Rand()
	model.Id = id.Hex()
	return Orm.Insert(model)
}
