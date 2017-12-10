package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	_ "testingWebApp/routers"
)

func main() {
	sessionconf := &session.ManagerConfig{
		CookieName: "begoosessionID",
		Gclifetime: 3600,
	}
	beego.GlobalSessions, _ = session.NewManager("memory", sessionconf)
	go beego.GlobalSessions.GC()

	beego.Run()
}
