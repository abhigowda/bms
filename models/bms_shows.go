package model

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type BmsShows struct {
	Id          int       `orm:"column(id);auto"`
	MovieId     int       `orm:"column(movieId);size(64);null" json:"movieId"`
	TheaterId   int       `orm:"column(theaterId);size(64);null" json:"theaterId"`
	ShowDate    time.Time `orm:"column(showDate);size(64);null" json:"showDate"`
	ShowPattern string    `orm:"column(showPattern);size(64);null" json:"showPattern"`
	CreatedOn   time.Time `orm:"column(createdOn);type(datetime);null;auto_now" json:"createdOn"`
	UpdatedOn   time.Time `orm:"column(updatedOn);type(datetime);null;auto_now"`
}

func (t *BmsShows) TableName() string {
	return "bms_shows"
}

func init() {
	orm.RegisterModel(new(BmsShows))
}
