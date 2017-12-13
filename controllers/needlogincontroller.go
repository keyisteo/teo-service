package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
)

type NeedLoginController struct {
	beego.Controller
	Session session.Store
}

func (this *NeedLoginController) Prepare() {
	this.Session = this.StartSession()
	this.Layout = "layout.tpl"
}
