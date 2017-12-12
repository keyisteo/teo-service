package controllers

import (
	"fmt"
	"strconv"
	"teo-service/models"
	"teo-service/utils/resize"
)

type ReportController struct {
	ExtendedController // Inherit our, slightly extended controller
}

func (c *ReportController) Get() {
	//Pass some data to the template
	c.Data["Title"] = "Report"
	c.TplName = "reportCreate.tpl"
}

func (c *ReportController) Post() {
	r := models.MyReport{}
	if err := c.ParseForm(&r); err != nil {
		//handle error
	}
	r.IdReporter = 1
	r.Id = models.InsertReport(r)
	file, _, _ := c.GetFile("picture")
	if file != nil {
		// get the filename
		var fileName string
		fileName = "static/img/reportImage/report_" + strconv.Itoa(r.Id) + ".jpg"
		// save to server
		c.SaveToFile("file", fileName)
		fmt.Println("Melewati penyimpanan file")
		//resize
		resize.ResizeImage(fileName, fileName)
		r.LinkPhoto = fileName
	}
	models.UpdateReport(r.Id, r.LinkPhoto)

}
