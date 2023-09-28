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
	flag.Parse()
	// 创建一个新的命令集，支持子命令
	goCmd := flag.NewFlagSet("go", flag.ExitOnError)
	goCmd.StringVar(&name, "name", "Go语言", "帮助信息")
	phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
	phpCmd.StringVar(&name, "n", "PHP语言", "帮助信息")

	args := flag.Args()
	switch args[0] {
	case "go":
		_ = goCmd.Parse(args[1:])
	case "php":
		_ = phpCmd.Parse(args[1:])

	}

	log.Printf("name: %s", name)

}
