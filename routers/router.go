package routers

import (
	"github.com/astaxie/beego"
	"teo-service/controllers"
)

func init() {
	//Home
	beego.Router("/", &controllers.MainController{})
	//Reporting
	beego.Router("/report", &controllers.ReportController{})
	//Authentication
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LogoutController{})
	beego.Router("/register", &controllers.RegisterController{})
	//
}
