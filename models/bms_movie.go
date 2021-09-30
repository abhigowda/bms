package model

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type BmsMovies struct {
	Id          int       `orm:"column(id);auto"`
	Name        string    `orm:"column(name);size(64);null" json:"name"`
	Description string    `orm:"column(description);size(64);null" json:"description"`
	Ratings     string    `orm:"column(ratings);size(64);null" json:"ratings"`
	Language    string    `orm:"column(language);size(64);null" json:"language"`
	Photo       string    `orm:"column(photo);size(13);null" json:"photo"`
	Genere      string    `orm:"column(genere);size(12);null" json:"genere"`
	CreatedOn   time.Time `orm:"column(createdOn);type(datetime);null;auto_now" json:"createdOn"`
	UpdatedOn   time.Time `orm:"column(updatedOn);type(datetime);null;auto_now"`
}

func (t *BmsMovies) TableName() string {
	return "bms_movie"
}

func init() {
	orm.RegisterModel(new(BmsMovies))
}

// GetBmsMoviesById retrieves BmsMovies by Id. Returns error if
// Id doesn't exist
func GetBmsMoviesById(id int) (v *BmsMovies, err error) {
	o := orm.NewOrm()
	return GetBmsMoviesByIdWithORM(id, o)
}

// GetBmsMoviesByIdWithORM retrieves BmsMovies by Id. Returns error if
// Id doesn't exist
func GetBmsMoviesByIdWithORM(id int, o orm.Ormer) (v *BmsMovies, err error) {
	v = &BmsMovies{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	fmt.Println(err)
	return nil, err
}

func GetBmsMoviesByEmail() (v *[]BmsMovies, err error) {
	o := orm.NewOrm()
	v = &[]BmsMovies{}
	_, err = o.QueryTable(new(BmsMovies)).All(v)
	if err == nil {
		return v, nil
	}
	return nil, err
}
