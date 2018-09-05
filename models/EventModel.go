package models

import (
	"web-pms/util"
)

type Event struct {
	Id            string `json:"id" xorm:"pk varchar(50)"`
	Version       string `json:"version" xorm:"varchar(50)"`
	ApiKey        string `json:"apikey" xorm:"varchar(255)"`
	DN            string `json:"dm" xorm:"varchar(255)"`
	UserAgent     string `json:"userAgent" xorm:"varchar(255)"`
	CookieEnabled bool   `json:"cookieEnabled" xorm:"default 'true' ENUM('true','false')"`
	//JavaEnabled   bool   `json:"javaEnabled" xorm:"default 'false' ENUM('true','false')"`
	Language     string `json:"language" xorm:"varchar(50)"`
	Url          string `json:"url" xorm:"varchar(255)"`
	Title        string `json:"title" xorm:"varchar(255)"`
	Referrer     string `json:"referrer" xorm:"default '' varchar(255)"`
	EventType    string `json:"type" xorm:"varchar(50)"`
	EventTime    int    `json:"time" xorm:"default 0 bigint(50)"`
	MetaData     string `json:"metaData" xorm:"varchar(255)"`
	Name         string `json:"name" xorm:"varchar(255)"`
	Message      string `json:"message" xorm:"varchar(255)"`
	FileName     string `json:"fileName" xorm:"varchar(255)"`
	LineNumber   int    `json:"lineNumber" xorm:"not null default 0 int"`
	ColumnNumber int    `json:"columnNumber" xorm:"not null default 0 int"`
	Stacktrace   string `json:"stacktrace" xorm:"varchar(255)"`
	Severity     string `json:"severity" xorm:"varchar(255)"`

	Target      Target `json:"target" xorm:"-"`
	Req         Req `json:"req" xorm:"-"`
	Res         Res `json:"res" xorm:"-"`
	Breadcrumbs interface{} `json:"breadcrumbs" xorm:"-"`
	Resources   []Resource  `json:"resources" xorm:"-"`
	PageTimes   PageTimes   `json:"pageTimes" xorm:"-"`
}

func InsertEvent(model *Event) (int64, error) {
	var id util.UUID = util.Rand()
	model.Id = id.Hex()
	return Orm.Insert(model)
}