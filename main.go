package main

import (
	"log"
	"os"
	"strings"

	"flagTest/controller"
)

func main() {

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
