package main

import (
	_ "Code_1218/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	fmt.Println("mytest")
	fmt.Println("hello world")
	beego.Run()


