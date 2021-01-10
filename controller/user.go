package controller

import (
	"fmt"
)

type UserController struct {
}

// 展示用户
func (this *UserController) List() {
	view = "index_view"
	users := userService.GetList()
	fmt.Println("| username  | password | age | sex |")
	for _, user := range users {
		fmt.Printf("| %s | %s | %d | %s |\n", user.GetUsername(), user.GetPassword(), user.GetAge(), user.GetSex())
	}
}
func (this *UserController) Update() {

}
