package controllers

import "github.com/astaxie/beego"

type VariableController struct {
	beego.Controller
}

var (
	//booléan
	a bool = true
	b bool = false

	//string
	c string = "Je suis un string"

	//integer
	d uint8 = 1
	e int8  = -1
	f       = float32(-1.22)
	g       = float32(1.22)

	//array
	h = [2]int8{1, 2}
	i = [2]string{"Hello", "World"}
	j = [2]float32{1.0, 2.0}
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
	v.Data["h"] = h
	v.Data["i"] = i
	v.Data["j"] = j

	v.LayoutSections = make(map[string]string)
	v.LayoutSections["HtmlHead"] = "head.html"
	v.LayoutSections["Header"] = "header.html"
	v.LayoutSections["Section"] = "section.html"
	v.LayoutSections["Footer"] = "footer.html"
}
