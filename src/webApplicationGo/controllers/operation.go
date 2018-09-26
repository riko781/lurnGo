package controllers

import (
	"github.com/astaxie/beego"
)

type OperationController struct {
	beego.Controller
}

func (c *OperationController) Get() {
	var x = 42
	var y = 30

	var resultAddition = x + y
	var resultMultiplication = x * y
	var resultDivision = x / y
	var resultSoustraction = (x - y)

	c.Data["x"] = x
	c.Data["y"] = y
	c.Data["resultAdittion"] = resultAddition
	c.Data["resultMultiplication"] = resultMultiplication
	c.Data["resultDivision"] = resultDivision
	c.Data["resultSoustraction"] = resultSoustraction
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "henrickcarneva@gmail.com"
	c.Data["DocGo"] = "golang.org"
	c.TplName = "site.tpl"
}
