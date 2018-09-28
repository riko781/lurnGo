package controllers

import (
	"runtime"

	"github.com/astaxie/beego"
)

type ConditionController struct {
	beego.Controller
}

func (c *ConditionController) Get() {
	c.Layout = "condition.tpl"
	c.TplName = "index.html"

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = "head.html"
	c.LayoutSections["Header"] = "header.html"
	c.LayoutSections["Section"] = "section.html"
	c.LayoutSections["Footer"] = "footer.html"
	c.Data["x"] = x
	c.Data["y"] = y

	//if condition
	if x < y {
		c.Data["a"] = a
	} else {
		c.Data["b"] = b
	}

	if x > y {
		c.Data["a"] = a
	} else {
		c.Data["b"] = b
	}

	if x := x + y; x > y {
		c.Data["v"] = a
	} else {
		c.Data["v"] = b
	}

	//switch condition
	switch os := runtime.GOOS; os {
	case "darwin":
		c.Data["os"] = "OS X."
	case "linux":
		c.Data["os"] = "Linux."
	default:
		c.Data["os"] = "Other OS"
	}
}
