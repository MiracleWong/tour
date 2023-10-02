package main

import (
	"flag"
	"log"
)

var name string

func main() {
	//var name string
	//flag.StringVar(&name, "name", "Go语言编程之旅", "帮助信息")
	//flag.StringVar(&name, "n", "Go语言编程之旅", "帮助信息")
	//flag.Parse()
	//
	//log.Printf("name: %s ", name)

	// 命令行定义为解析的标志
	// 解析并绑定命令行的参数
	flag.Parse()

	// 创建一个新的命令集，NewFlagSet支持子命令
	// 方法第二个参数是ErrorHandling，有ContinueOnError、ExitOnError、PanicOnError三种模式
	goCmd := flag.NewFlagSet("go", flag.ExitOnError)
	goCmd.StringVar(&name, "name", "Go语言", "帮助信息")
	phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
	phpCmd.StringVar(&name, "n", "PHP语言", "帮助信息")

	// Args把新的空命令集做为外部参数传入
	args := flag.Args()
	switch args[0] {
	case "go":
		// 对解析方法的进一步封装
		_ = goCmd.Parse(args[1:])
	case "php":
		_ = phpCmd.Parse(args[1:])

	}

	log.Printf("name: %s", name)

}
