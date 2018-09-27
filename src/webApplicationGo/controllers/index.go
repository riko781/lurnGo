package controllers

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	c.Layout = "acceuil.tpl"
	c.TplName = "index.html"

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = "head.html"
	c.LayoutSections["Header"] = "header.html"
	c.LayoutSections["Section"] = "section.html"
	c.LayoutSections["Footer"] = "footer.html"
}
