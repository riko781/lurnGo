package controllers

import "github.com/astaxie/beego"

type BoucleController struct {
	beego.Controller
}

var (
	somme [10]int
	slice = []string{"henrick", "melina", "ericka"}
)

func (b *BoucleController) Get() {

	b.Layout = "boucle.tpl"
	b.TplName = "index.html"
	b.LayoutSections = make(map[string]string)
	b.LayoutSections["HtmlHead"] = "head.html"
	b.LayoutSections["Header"] = "header.html"
	b.LayoutSections["Section"] = "section.html"
	b.LayoutSections["Footer"] = "footer.html"

	for i := 0; i < 10; i++ {
		somme[i] = i
		b.Data["somme"] = somme

	}

	b.Data["slice"] = slice
}
