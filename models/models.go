package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type MyReport struct {
	Id                 int       `form:"-"`
	IdReporter         int       `form:"-"`
	VerificationStatus bool      `form:"-"`
	HandleReport       int       `form:"-"`
	Category           int       `form:"category"`
	Detail             string    `form:"detail"`
	LinkPhoto          string    `form:"detail"`
	CreatedAt          time.Time `form:"-"`
	UpdatedAt          time.Time `form:"-"`
}

func InsertReport(r MyReport) int {
	o := orm.NewOrm()
	res, err := o.Raw("INSERT INTO reports (id_reporter, category, detail, link_photo,created_at,updated_at) VALUES(?,?,?,?,?,?)", r.IdReporter, r.Category, r.Detail, r.LinkPhoto, time.Now(), time.Now()).Exec()
	if err == nil {
		ret, _ := res.LastInsertId()
		fmt.Println("mysql row affected nums: ", ret)
		return int(ret)
	}

	return 0
}

func UpdateReport(id int, s string) {
	o := orm.NewOrm()
	o.Raw("UPDATE reports SET link_photo= ?, updated_at=? WHERE id= ?", s, time.Now(), id).Exec()
}
