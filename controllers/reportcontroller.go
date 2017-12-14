package controllers

import (
	"fmt"
	"html/template"
	"strconv"
	"teo-service/models"
	"teo-service/utils/resize"
)

type ReportController struct {
	ExtendedController // Inherit our, slightly extended controller
}

func (c *ReportController) Get() {
	userDataIn := c.Session.Get("UserData")
	if userDataIn != nil {
		userData := userDataIn.(models.User)
		fmt.Println(userData.Username)
		c.Data["userData"] = userData
		id := 0
		c.Ctx.Input.Bind(&id, "id")
		if id == 0 { //Pass some data to the template
			c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
			c.Data["Title"] = "Report"
			c.TplName = "reportCreate.tpl"
		} else {
			r := models.MyReport{}
			r = models.RetrieveOneReport(id)
			if r.Id == 0 {
				c.Ctx.Redirect(302, "/report")
			} else {
				c.Show(r.Id)
			}
		}
	} else {
		c.Ctx.Redirect(302, "/")
	}
}

func (c *ReportController) Post() {
	r := models.MyReport{}
	if err := c.ParseForm(&r); err != nil {
		//handle error
		c.Ctx.Redirect(302, "/report")
		return
	}
	r.IdReporter = 1
	r.Id = models.InsertReport(r)
	file, _, _ := c.GetFile("picture")

	if file != nil {
		// get the filename
		var fileName string
		fileName = "static/img/reportImage/report_" + strconv.Itoa(r.Id) + ".jpg"
		// save to server
		c.SaveToFile("picture", fileName)
		fmt.Println("Melewati penyimpanan file")
		//resize
		resize.ResizeImage(fileName, fileName)
		r.LinkPhoto = fileName
	}
	models.UpdateReport(r.Id, r.LinkPhoto)

	c.Ctx.Redirect(302, "/report?id="+strconv.Itoa(r.Id))
}

// @router /report/:id
func (c *ReportController) Show(id int) {
	r := models.RetrieveOneReport(id)
	p := models.RetriveNameBasedID(r.IdReporter)
	c.Data["Title"] = "Show Report"
	c.Data["Report"] = r
	c.Data["Pelapor"] = p
	c.TplName = "reportShow.tpl"
}

func (c *ReportController) Delete() {
	id := 0
	c.Ctx.Input.Bind(&id, "id")
	if id != 0 {
		//Melakukan delet
		fmt.Println("akan melakukan delete")
		models.DeleteReport(id)
		return
	}
	c.Ctx.Redirect(302, "/timeline")
	return
}
