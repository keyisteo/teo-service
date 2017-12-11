package controllers

import (
	"github.com/astaxie/beego"
)

type ExtendedController struct {
	beego.Controller //inherit
}

func (this *ExtendedController) Prepare() {
	this.Layout = "layout.tpl"
}
