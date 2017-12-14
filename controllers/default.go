package controllers

import (
	"html/template"
)

type MainController struct {
	ExtendedController
}

// @Title getMainController
// @Summary Show Homepage
// @Description get all the staticblock by key
// @Success 200 OK
// @Failure 400 Bad request
// @Failure 404 Not found
// @router / [get]

func (c *MainController) Get() {
	// Check if user is logged in
	userData := c.Session.Get("UserData")

	// Do input checks
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["Title"] = "Selamat Datang"
	c.Data["userData"] = userData
	c.TplName = "welcome.tpl"
}
