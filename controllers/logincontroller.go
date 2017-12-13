package controllers

import (
	"html/template"
	"teo-service/models"
)

type LoginController struct {
	ExtendedController // Inherit our, slightly extended controller
}

func (c *LoginController) Get() {
	// Check if user is logged in
	session := c.StartSession()
	userData := session.Get("UserData")

	if userData != nil {
		// User is logged in already, display another page
		c.Ctx.Redirect(302, "/report")
		return
	}

	// Do input checks
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["Title"] = "Login"
	c.TplName = "login.tpl"
}

func (c *LoginController) Post() {
	// Check if user is logged in
	session := c.StartSession()
	userData := session.Get("UserData")

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

	user = models.RetrieveOneUser(user.Username, user.Password)
	if user.Id == 0 {
		c.Ctx.Redirect(302, "/report")
		return
	} else {
		// Set the userData if everything is ok
		session.Set("UserData", user)
		c.Data["json"] = user
		c.ServeJSON()

	}

}
