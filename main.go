package main

import (
	_ "Blog/models"
	_ "Blog/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
