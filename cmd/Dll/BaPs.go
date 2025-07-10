package main

import (
	"log"

	"./"
)

func main() {
	log.Print("BaPs动态库加载成功")
}

//export StartServer
func StartServer() {
	BaPs.NewBaPs()
}
