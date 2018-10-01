package controllers

import "github.com/astaxie/beego"

type FunctionController struct {
	beego.Controller
}

var (
	nom    string
	prenom string
)

func name(nom string) string {
	return nom
}

func nameMultiple(nom, prenom string) (string, string) {
	return nom, prenom
}

func (f *FunctionController) Get() {

	f.Layout = "function.tpl"
	f.TplName = "index.html"
	f.LayoutSections = make(map[string]string)
	f.LayoutSections["HtmlHead"] = "head.html"
	f.LayoutSections["Header"] = "header.html"
	f.LayoutSections["Section"] = "section.html"
	f.LayoutSections["Footer"] = "footer.html"

	f.Data["name"] = name("Carneva")
	nom, prenom = nameMultiple("Carneva", "henrick")
	f.Data["nom"] = nom
	f.Data["prenom"] = prenom
}
