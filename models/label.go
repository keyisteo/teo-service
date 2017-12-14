package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Label struct {
	Id       int    `json:"id"`
	Category string `json:"label"`
}

func JSONRetriveLabelById(id int) Label {
	o := orm.NewOrm()
	l := Label{}
	_ = o.Raw("SELECT id,category FROM label_reports WHERE id=?", id).QueryRow(&l)
	fmt.Println(l.Id)
	return l
}

func JSONRetriveLabelAll() []Label {
	o := orm.NewOrm()
	var l []Label
	_, _ = o.Raw("SELECT id,category FROM label_reports").QueryRows(&l)
	return l
}
