package controllers

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "henrickcarneva@gmail.com"
	c.Data["DocGo"] = "golang.org"
	c.TplName = "site.tpl"

}
