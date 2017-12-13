package routers

import (
	"github.com/astaxie/beego"
	"teo-service/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/report", &controllers.ReportController{})
	beego.Router("/login", &controllers.LoginController{})
}
