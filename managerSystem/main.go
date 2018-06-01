package main

import (
	"github.com/astaxie/beego"
	"managerSystem/models"
	_ "managerSystem/routers"
	"mime"
	"strconv"
)

func StringsToJson(str string) string {
	rs := []rune(str)
	jsons := ""
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			jsons += string(r)
		} else {
			jsons += "\\u" + strconv.FormatInt(int64(rint), 16) // json
		}
	}

	return jsons
}

func main() {

	models.Syncdb()
	mime.AddExtensionType(".css", "text/css")

	beego.AddFuncMap("stringsToJson", StringsToJson)
	beego.Run()
}

//func Run() {
//	//初始化
//	initialize()
//
//	fmt.Println("Starting....")
//
//	fmt.Println("Start ok")
//}
//func initialize() {
//	mime.AddExtensionType(".css", "text/css")
//	//判断初始化参数
//	initArgs()
//
//	models.Connect()
//
//	router()
//	beego.AddFuncMap("stringsToJson", StringsToJson)
//}
//func initArgs() {
//	args := os.Args
//	for _, v := range args {
//		if v == "-syncdb" {
//			models.Syncdb()
//			os.Exit(0)
//		}
//	}
//}
