package main

import (
	_ "blog/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("tablename", "mysql", "root:root@tcp(localhost:8889)/blog_bd")
	beego.Run()
}
