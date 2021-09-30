package model

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/spf13/cast"
)

type BmsBookings struct {
	Id         int       `orm:"column(id);auto"`
	UserId     int       `orm:"column(userId);size(64);null" json:"userId"`
	MovieId    int       `orm:"column(movieId);size(64);null" json:"movieId"`
	TheaterId  int       `orm:"column(theaterId);size(64);null" json:"theaterId"`
	ShowDate   time.Time `orm:"column(showDate);size(64);null" json:"showDate"`
	ShowTime   string    `orm:"column(showTime);size(64);null" json:"showTime"`
	TotalSeats int       `orm:"column(totalSeats);size(13);null" json:"totalSeats"`
	CreatedOn  time.Time `orm:"column(createdOn);type(datetime);null;auto_now" json:"createdOn"`
	UpdatedOn  time.Time `orm:"column(updatedOn);type(datetime);null;auto_now"`
}

func (t *BmsBookings) TableName() string {
	return "bms_bookings"
}

func init() {
	orm.RegisterModel(new(BmsBookings))
}

// AddBmsBookings insert a new BmsBookings into database and returns
// last inserted Id on success.
func AddBmsBookings(m *BmsBookings) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	fmt.Println(err)
	return
}

func GetTotalBooked(showDate, showTime string, theaterId, movieId int) (int, error) {
	o := orm.NewOrm()
	var v []orm.Params
	_, err := o.Raw(`SELECT SUM(totalSeats) AS seats FROM bms_bookings WHERE movieId = ? AND theaterId = ? AND showDate = ? AND showTime = ?`, movieId, theaterId, showDate, showTime).Values(&v)
	if err == nil && len(v) > 0 {
		return cast.ToInt(v[0]["seats"]), nil
	}
	fmt.Println(err)
	return 0, err
}

func GetBookingBYuserId(id int) (v *[]BmsBookings, err error) {
	o := orm.NewOrm()
	v = &[]BmsBookings{}
	if _, err = o.QueryTable(new(BmsBookings)).Filter("userId", id).RelatedSel().All(v); err == nil {
		return v, nil
	}
	fmt.Println(err)
	return nil, err
}
