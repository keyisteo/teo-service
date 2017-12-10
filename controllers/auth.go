package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

type LogoutController struct {
	beego.Controller
}

type UserLoginRequiredController struct {
	beego.Controller
}

func (this *LoginController) Post() {
	// Check if user is logged in
	session := this.StartSession()
	userID := session.Get("UserID")

	if userID != nil {
		// User is logged in already, display another page
		return
	}

	// Do input checks

	// Set the UserID if everything is ok
	session.Set("UserID", 50)
}

func (this *LogoutController) Get() {
	// Check if user is logged in
	session := this.StartSession()
	userID := session.Get("UserID")
	if userID != nil {
		// UserID is set and can be deleted
		session.Delete("UserID")
	}
}

func (this *UserLoginRequiredController) Get() {
	// Check if user is logged in
	session := this.StartSession()
	userID := session.Get("UserID")

	if userID == nil {
		// UserID is not set, display another page
		return
	}

	// Display page that requires to be logged in
}
