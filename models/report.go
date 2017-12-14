package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type MyReport struct {
	Id         int       `form:"-" json:"id"`
	IdReporter int       `form:"-" json:"id_reporter"`
	Category   int       `form:"category" json:"category"`
	Detail     string    `form:"detail" json:"detail"`
	LinkPhoto  string    `form:"link_photo" json:"link_photo"`
	CreatedAt  time.Time `form:"-" json:"create_time"`
	UpdatedAt  time.Time `form:"-" json:"-"`
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(MyReport))
}

func (r *MyReport) TableName() string {
	return "reports"
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

func RetrieveOneReport(id int) MyReport {
	o := orm.NewOrm()
	report := MyReport{Id: id}
	err := o.Read(&report)
	if err == orm.ErrNoRows {
		fmt.Println("No result found.")
		report.Id = 0
		return report
	} else if err == orm.ErrMissPK {
		fmt.Println("No primary key found.")
		report.Id = 0
		return report
	} else {
		return report
	}
}

func JSONRetrieveReportById(id int) MyReport {
	o := orm.NewOrm()
	r := MyReport{Id: id}
	_ = o.Read(&r)
	return r
}

func JSONRetrieveReportAll() []MyReport {
	o := orm.NewOrm()
	var r []MyReport
	_, _ = o.Raw("SELECT * FROM reports").QueryRows(&r)
	return r
}

func DeleteReport(id int) {
	o := orm.NewOrm()
	r := MyReport{Id: id}
	o.Delete(&r)
}
