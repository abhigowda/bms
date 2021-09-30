package model

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type BmsUser struct {
	Id        int       `orm:"column(id);auto"`
	Name      string    `orm:"column(name);size(64);null" json:"name"`
	AuthId    string    `orm:"column(authId);size(64);null" json:"authId"`
	Mobile    string    `orm:"column(mobile);size(64);null" json:"mobile"`
	Email     string    `orm:"column(Email);size(64);null" json:"Email"`
	Password  string    `orm:"column(password);size(64);null" json:"password"`
	Gender    string    `orm:"column(Gender);size(13);null" json:"Gender"`
	CreatedOn time.Time `orm:"column(createdOn);type(datetime);null;auto_now" json:"createdOn"`
	UpdatedOn time.Time `orm:"column(updatedOn);type(datetime);null;auto_now"`
}

func (t *BmsUser) TableName() string {
	return "bms_user"
}

func init() {
	orm.RegisterModel(new(BmsUser))
}

// AddBmsUser insert a new BmsUser into database and returns
// last inserted Id on success.
func AddBmsUser(m *BmsUser) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	fmt.Println(err)
	return
}

func GetBmsUserByEmail(email string) (v *BmsUser, err error) {
	o := orm.NewOrm()
	v = &BmsUser{}
	err = o.QueryTable(new(BmsUser)).Filter("email", email).One(v)
	if err == nil {
		return v, nil
	}
	return nil, err
}
