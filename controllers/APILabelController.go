package controllers

import (
	"teo-service/models"
)

type LabelController struct {
	ExtendedController // Inherit our, slightly extended controller
}

func (c *LabelController) Get() {
	id := 0
	c.Ctx.Input.Bind(&id, "id")
	if id == 0 { //Pass some data to the template
		json := models.JSONRetriveLabelAll()
		c.Data["json"] = &json
		c.ServeJSON()
	} else {
		c.Show(id)
	}

}

// @router /report/:id
func (c *LabelController) Show(id int) {
	json := models.JSONRetriveLabelById(id)
	c.Data["json"] = &json
	c.ServeJSON()
}
