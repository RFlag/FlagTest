package controller

import (
	"flagTest/conf"
	"flagTest/model"
	"log"
	"net/http"
	"os"
)

func TestData(paramMap map[string]string) {
	okAll := true
	for group, urlPrefix := range paramMap {
		for _, n := range conf.Conf[group] {
			req, err := http.NewRequest(n.Method, urlPrefix+n.Url, nil)
			if err != nil {
				log.Println(ResultHttpResuestCreateFailed, err)
				return
			}
			for head, _ := range n.Header {
				for _, token := range n.Header[head] {
					req.Header.Set(head, token)
				}
			}
			outPut, ok, err := model.TestData(req, n.Result)
			if err != nil {
				log.Println(err)
				continue
			}

			if ok == false {
				log.Printf("url: %v\nreturnResult: %v\nresult: filed\n", urlPrefix+n.Url, outPut)
				okAll = false
			} else {
				log.Printf("url: %v\nresult: pass\n", urlPrefix+n.Url)
			}
		}
	}

	if !okAll {
		os.Exit(1)
		return
	}
}
