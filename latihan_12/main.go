package main

import (
	"authman/permission"
	"flag"
)

func main() {

	var namefile string
	var permission_binary string

	flag.StringVar(&namefile, "f", "", "help")
	flag.StringVar(&permission_binary, "p", "", "help")

	permission.Run(namefile, permission_binary)

}
