package controller

import "fmt"

type IndexController struct {
}

func (this *IndexController) Index() {
	view = "index_view"
	fmt.Println("这里是首页，没有什么信息")
}

func (this *IndexController) Welcome() {
	fmt.Println("欢迎来到 六星教育学员管理系统 cmdedu")
	fmt.Println("你的执行操作：")
	view = "login_view"
}
