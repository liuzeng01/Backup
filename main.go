package main

import (
	_ "GoStudy/week5/backup/routers"
	ui "GoStudy/week5/backup/uilogic"

	"github.com/astaxie/beego"
)

func main() {
	go beego.Run()

	ui.LoadUi()

}
