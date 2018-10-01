package controllers

import "github.com/astaxie/beego"

type PointerController struct {
	beego.Controller
}

var ()

func (p *PointerController) Get() {

	p.Layout = "pointer.tpl"
	p.TplName = "index.html"
	p.LayoutSections = make(map[string]string)
	p.LayoutSections["HtmlHead"] = "head.html"
	p.LayoutSections["Header"] = "header.html"
	p.LayoutSections["Section"] = "section.html"
	p.LayoutSections["Footer"] = "footer.html"

	i, j := 42, 2701
	po := &i
	p.Data["b"] = *po
	po = &j
	*po = (*po / 37)
	p.Data["j"] = j
}
