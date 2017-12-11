package controllers

type ReportController struct {
	ExtendedController // Inherit our, slightly extended controller
}

func (c *ReportController) Get() {
	//Pass some data to the template
	c.Data["Title"] = "Report!"
	c.TplName = "reportCreate.tpl"
}
