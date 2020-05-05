package ui

import (
	"fmt"
	"io"
	"log"
	"os"
)

func cronexpression(cronstring string) {
	fmt.Println("cronexpression", cronstring)
	cronstore(cronstring)

}

func cronstore(cronstring string) {
	file, err := os.OpenFile("cron.ini", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	file.Write([]byte(cronstring)) //写入字节切片数据
}

func configitemsSave(localitems map[string]string, sftpitems map[string]string, Encryptitems map[string]string) {
	var item configitem
	item.localitems = localitems
	item.sftpitems = sftpitems
	item.Encryptitems = Encryptitems
	ParseText(item)
}

func buildpackage() {
	os.MkdirAll("./build", os.ModePerm)
	copyFile("./conf/build", "./build/build")
	copyFile("gobackup.yaml", "./build/gobackup.yaml")
	copyFile("cron.ini", "./build/cron.ini")
}
func copyFile(srcFile, destFile string) {
	file1, err := os.Open(srcFile)
	if err != nil {
		log.Println(err)
	}
	file2, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		log.Println(err)
	}
	defer file1.Close()
	defer file2.Close()
	_, err = io.Copy(file2, file1)
	log.Println(err)
	fmt.Printf("%s copy successfully!!!", srcFile)
}
