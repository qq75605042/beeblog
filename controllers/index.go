package controllers

import (
	"beeproject/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Layout = "layout/layout.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["CustomCss"] = ""
	c.LayoutSections["Nav"] = "layout/nav.html"
	c.LayoutSections["Footer"] = "layout/footer.html"
	c.LayoutSections["Scripts"] = ""
	arts,err := models.GetAllArticle(nil,nil,[]string{"id"},[]string{"asc"},0,10)
	if err != nil {
		//log.Fatal(err)
	}
	//log.Printf("%v",arts)
	c.Data["articles"] = arts
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}
