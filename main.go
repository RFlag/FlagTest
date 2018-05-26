package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"flagTest/controller"
)

func main() {
	if len(os.Args) <= 1 {
		help()
		return
	}
	paramMap := map[string]string{}
	if len(os.Args) > 1 {
		for _, v := range os.Args[1:] {
			n := strings.Index(v, "=")
			if n == -1 {
				log.Println(controller.ResultParamError)
				os.Exit(2)
			}
			paramMap[v[:n]] = v[n+1:]
		}
		controller.TestData(paramMap)
	}
}

func help() {
	fmt.Println("语法: flagTest [参数]")
	fmt.Println()
	fmt.Println("go 参数:serverName=serverUrl")
	fmt.Println("\tserverName:服务名\t")
	fmt.Println("\tserverUrl:服务url\t")
	fmt.Println("例:\tflagTest webapp=http://api.private wallet=http://localhost:8080")
}
