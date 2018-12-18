package controllers

import (
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
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}
