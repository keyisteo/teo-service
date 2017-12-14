package controllers

import (
	"teo-service/models"
)

type TimelineController struct {
	ExtendedController // Inherit our, slightly extended controller
}

func (c *TimelineController) Get() {

	userData := c.Session.Get("UserData")
	if userData == nil {
		// User is logged in already, display another page
		c.Ctx.Redirect(302, "/report")
		return
	}
	r := models.JSONRetrieveReportAll()
	u := models.RetrieveAllUser()
	c.Data["Title"] = "Linimasa Report"
	c.Data["Report"] = &r
	c.Data["Pelapor"] = &u
	c.Data["userData"] = &userData
	c.TplName = "reportIndex.tpl"

}
