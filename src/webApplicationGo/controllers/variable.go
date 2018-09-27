package controllers

import "github.com/astaxie/beego"

type VariableController struct {
	beego.Controller
}

var (
	a bool   = true
	b bool   = false
	c string = "Je suis un string"
	d uint8  = 1
	e int8   = -1
	f        = float32(-1.22)
	g        = float32(1.22)
)

func (v *VariableController) Get() {
	v.Layout = "variable.tpl"
	v.TplName = "index.html"

	v.Data["a"] = a
	v.Data["b"] = b
	v.Data["c"] = c
	v.Data["d"] = d
	v.Data["e"] = e
	v.Data["f"] = f
	v.Data["g"] = g

	v.LayoutSections = make(map[string]string)
	v.LayoutSections["HtmlHead"] = "head.html"
	v.LayoutSections["Header"] = "header.html"
	v.LayoutSections["Section"] = "section.html"
	v.LayoutSections["Footer"] = "footer.html"
}
