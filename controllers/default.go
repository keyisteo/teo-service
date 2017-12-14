package controllers

import (
	"html/template"
)

type MainController struct {
	ExtendedController
}

func (c *MainController) Get() {
	// Check if user is logged in
	userData := c.Session.Get("UserData")

	// Do input checks
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["Title"] = "Selamat Datang"
	c.Data["userData"] = userData
	c.TplName = "welcome.tpl"
}
