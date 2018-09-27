package controllers

import (
	"github.com/astaxie/beego"
)

type OperationController struct {
	beego.Controller
}

var (
	x uint16 = 42
	y uint16 = 30

	resultAddition       uint16  = x + y
	resultMultiplication uint16  = x * y
	resultDivision       float32 = float32(x) / float32(y)
	resultSoustraction   uint16  = x - y
)

func (c *OperationController) Get() {
	c.Layout = "operation.tpl"
	c.TplName = "index.html"
	
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = "head.html"
	c.LayoutSections["Header"] = "header.html"
	c.LayoutSections["Section"] = "section.html"
	c.LayoutSections["Footer"] = "footer.html"

	c.Data["x"] = x
	c.Data["y"] = y
	c.Data["resultAdittion"] = resultAddition
	c.Data["resultMultiplication"] = resultMultiplication
	c.Data["resultDivision"] = resultDivision
	c.Data["resultSoustraction"] = resultSoustraction
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "henrickcarneva@gmail.com"
	c.Data["DocGo"] = "golang.org"

	
}