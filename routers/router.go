package routers

import (
	"DbHelpDreamEbagEfficientReading/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
