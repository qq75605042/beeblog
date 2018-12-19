package main

import (
	_ "beeproject/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	_ = orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/blog?charset=utf8&loc=Asia%2FShanghai", 30)
	// 开启调试模式,打印sql
	orm.Debug = true
	_ = logs.SetLogger(logs.AdapterFile,`{"filename":"sql.log"}`)
	//orm.RegisterModel(new(models.Article))
	//o:= orm.NewOrm()
	//art := models.Article{Id:int64(1)}
	//err := o.Read(&art)
	//fmt.Printf("%v",err)
	//fmt.Printf("%v",art)

	beego.Run()
}

