package model

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type BmsTheater struct {
	Id            int       `orm:"column(id);auto"`
	Name          string    `orm:"column(name);size(64);null" json:"name"`
	Address       string    `orm:"column(address);size(64);null" json:"address"`
	TotalCapacity int       `orm:"column(totalCapacity);size(64);null" json:"totalCapacity"`
	TotalShows    string    `orm:"column(totalShows);size(64);null" json:"totalShows"`
	CreatedOn     time.Time `orm:"column(createdOn);type(datetime);null;auto_now" json:"createdOn"`
	UpdatedOn     time.Time `orm:"column(updatedOn);type(datetime);null;auto_now"`
}

func (t *BmsTheater) TableName() string {
	return "bms_theater"
}

func init() {
	orm.RegisterModel(new(BmsTheater))
}

func GetBmsTheaterById(id int) (v *BmsTheater, err error) {
	o := orm.NewOrm()
	v = &BmsTheater{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	fmt.Println(err)
	return nil, err
}

func GetTheaterByDateAndMovieId(showDate string, movieId int) (v []orm.Params, err error) {
	o := orm.NewOrm()
	_, err = o.Raw(`SELECT bt.id AS theaterId, bt.name AS theaterName, bt.address, bt.totalCapacity, bs.showDate, bs.showPattern FROM bms_theater AS bt JOIN bms_shows bs ON bt.id = bs.theaterId WHERE bs.showDate = ? AND bs.movieId = ?`, showDate, movieId).Values(&v)
	if err == nil && len(v) > 0 {
		return v, nil
	}
	fmt.Println(err)
	return nil, err
}
