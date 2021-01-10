package util

import (
	"flag"
	"fmt"
)

var (
	instance *uconfig
	conf     = flag.String("conf", "../etc/config.json", "cmdedu . config")
)

type uconfig struct {
	BasePath string `json:"base_path"`
	DatePath string `json:"date_path"`
}

func init() {
	fmt.Println("uitl -->> ")
	flag.Parse()
	NewUConfigWithFile(*conf)
}

func NewUConfigWithFile(name string) *uconfig {
	if instance == nil {
		c := &uconfig{}
		ReadJosn(name, c)
		instance = c // <--- NOT THREAD SAFE
	}
	return instance
}
func GetConfig() *uconfig {
	return instance
}
func (u *uconfig) GetDatePath() string {
	return u.DatePath
}
