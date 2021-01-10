package main

import (
	"fmt"
	"cmdedu/controller"
)

func main() {
	fmt.Println("配置文件名：config.json")
	fmt.Println(`配置信息为json参考
		windows : {
	    "base_path" : "xxx",
	    "date_path" : "D:\\phpStudy\\PHPTutorial\\WWW\\go\\src\\cmdedu\\date\\"
	 };
	 linux:{
		 "base_path" : "/go/worker",
		 "date_path" : "/go/worker/data/"
	 }
	 `)
	fmt.Println("启动命令 edu -config=config.json")
	// util.GetInstance()
	// flag.Parse()

	// 1. 进入系统界面
	// fmt.Println(*configFile)
	controller.Run()
}
