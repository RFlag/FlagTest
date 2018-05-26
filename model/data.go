package model

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type aa struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	En   string `json:"en"`
}

func TestData(req *http.Request, r string) (string, bool, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", false, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", false, err
	}
	var d aa

	err = json.Unmarshal(body, &d)
	if err != nil {
		return "", false, err
	}
	l, _ := json.Marshal(d)

	result := string(l)
	if result == r {
		return "", true, nil
	}
	return result, false, nil
}
