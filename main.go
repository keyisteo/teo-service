package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/session"
	_ "github.com/go-sql-driver/mysql"
	_ "teo-service/routers"
	"time"
)

func init() {

	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:root@/kesat?charset=utf8")

	orm.DefaultTimeLoc = time.UTC

}

func main() {
	sessionconf := &session.ManagerConfig{
		CookieName: "begoosessionID",
		Gclifetime: 3600,
	}
	beego.GlobalSessions, _ = session.NewManager("memory", sessionconf)
	go beego.GlobalSessions.GC()

	beego.Run()
}
