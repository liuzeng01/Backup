package routers

import (
	"GoStudy/week5/backup/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/cron.html", &controllers.CronController{})
}
