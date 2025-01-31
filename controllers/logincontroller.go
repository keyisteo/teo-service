package controllers

import (
	//"html/template"
	"teo-service/models"
)

type LoginController struct {
	ExtendedController // Inherit our, slightly extended controller
}

// @Title getLogin
// @Summary Show Login Page
// @Description get html page of login
// @Success 200 OK
// @Failure 302 Already Login
// @router /login [get]

func (c *LoginController) Get() {
	// Check if user is logged in
	userData := c.Session.Get("UserData")

	if userData != nil {
		// User is logged in already, display another page
		c.Ctx.Redirect(302, "/report")
		return
	}
	// Do input checks
	//	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["Title"] = "Login"
	c.TplName = "login.tpl"
}

// @Title postLogin
// @Summary post data to login page
// @Description post data from form to process and authenticate user
// @Success 200 OK
// @Failure 302 Already Login
// @router /login [post]

func (c *LoginController) Post() {
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
	user = models.RetrieveOneUser(user.Username, user.Password)
	if user.Id == 0 {
		c.Ctx.Redirect(302, "/login")
		return
	} else {
		// Set the userData if everything is ok
		c.Session.Set("UserData", user)
		c.Ctx.Redirect(302, "/")
	}

}
