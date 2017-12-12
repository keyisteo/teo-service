package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
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
	beego.Run()
}
