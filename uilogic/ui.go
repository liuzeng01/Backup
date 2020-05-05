package ui

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/zserge/lorca"
)

func LoadUi() {
	ui, err := lorca.New("", "", 640, 480, "--disable-translate", "--disable-windows10-custom-titlebar", "--no-default-browser-check")
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()
	ui.Load("http://127.0.0.1:30450")

	Bindlogic(ui)

	// Wait until the interrupt signal arrives or browser window is closed
	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	case <-ui.Done():
	}
	Removefile("./cron.ini")
	Removefile("./gobackup.yaml")
	log.Println("exiting...")
}

func Bindlogic(ui lorca.UI) {
	ui.Bind("cronexpression", cronexpression)
	ui.Bind("configitemsSave", configitemsSave)
	ui.Bind("buildpackage", buildpackage)

}

func Removefile(filename string) {
	err := os.Remove(filename)
	if err != nil {
		fmt.Println(err)
	}
}
