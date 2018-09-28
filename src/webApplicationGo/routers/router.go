package routers

import (
	"webApplicationGo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/index", &controllers.IndexController{})
	beego.Router("/operation", &controllers.OperationController{})
	beego.Router("/variable", &controllers.VariableController{})
	beego.Router("/condition", &controllers.ConditionController{})
}
