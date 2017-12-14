package controllers

import (
	"html/template"
	"teo-service/models"
)

type RegisterController struct {
	ExtendedController // Inherit our, slightly extended controller
}

func (c *RegisterController) Get() {
	// Check if user is logged in
	userData := c.Session.Get("UserData")
	if userData != nil {
		// User is logged in already, display another page
		c.Ctx.Redirect(302, "/report")
		return
	}

	// Do input checks
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["Title"] = "Register"
	c.TplName = "register.tpl"
}

func (c *RegisterController) Post() {
	// Check if user is logged in
	userData := c.Session.Get("UserData")

	if userData != nil {
		// User is logged in already, display another page
		c.Ctx.Redirect(302, "/report")
		return
	}
	user := models.User{}
	if err := c.ParseForm(&user); err != nil {
		//handle error
	}
	// Do input checks
	checkError := models.InsertUser(user)
	if checkError == 0 {
		//Berarti error
	} else {
		//Berhsasil dan redirect ke halaman login
	}
	c.Ctx.Redirect(302, "/login")
}
