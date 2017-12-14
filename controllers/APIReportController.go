package controllers

import (
	"teo-service/models"
)

type APIReportController struct {
	ExtendedController // Inherit our, slightly extended controller
}

func (c *APIReportController) Get() {
	userData := c.Session.Get("UserData")
	if userData != nil {
		// User is logged in already, display another page
		c.Ctx.Redirect(302, "/report")
		return
	}
	id := 0
	c.Ctx.Input.Bind(&id, "id")
	if id == 0 { //Pass some data to the template
		json := models.JSONRetrieveReportAll()
		c.Data["json"] = &json
		c.ServeJSON()
	} else {
		c.Show(id)
	}

}

// @router /report/:id
func (c *APIReportController) Show(id int) {
	userData := c.Session.Get("UserData")
	if userData != nil {
		// User is logged in already, display another page
		c.Ctx.Redirect(302, "/report")
		return
	}
	json := models.JSONRetrieveReportById(id)
	c.Data["json"] = &json
	c.ServeJSON()
}
