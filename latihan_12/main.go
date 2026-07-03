package main

import (
	"authman/permission"
	"flag"
)

func main() {

	var namefile string
	var permission_binary string
	var show string

	flag.StringVar(&namefile, "f", "", "help")
	flag.StringVar(&show, "s", "", "help")
	flag.StringVar(&permission_binary, "p", "", "help")

	flag.Parse()
	permission.Run(namefile, permission_binary, show)
}
