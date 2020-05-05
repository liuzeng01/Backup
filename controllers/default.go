package controllers

import (
	"bufio"
	"io"
	"os"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	contents := Readfile("gobackup.yaml")
	c.Data["buildplan"] = contents
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}

type CronController struct {
	beego.Controller
}

func (c *CronController) Get() {
	c.TplName = "cron.html"
}

func Readfile(filename string) []string {
	var contents []string
	file, err := os.Open(filename)
	if err != nil {
		contents = append(contents, "You had not configure the backup plan!")
		return contents

	}
	defer file.Close()
	rd := bufio.NewReader(file)
	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		} else {
			contents = append(contents, line)
		}
	}
	return contents

}
