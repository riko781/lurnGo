package routers

import (
	"webApplicationGo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/beego", &controllers.MainController{})
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/operation", &controllers.OperationController{})
	beego.Router("/variable", &controllers.VariableController{})
	beego.Router("/condition", &controllers.ConditionController{})
	beego.Router("/boucle", &controllers.BoucleController{})
	beego.Router("/fonction", &controllers.FunctionController{})
}
