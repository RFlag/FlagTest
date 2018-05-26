package model

import (
	"io/ioutil"
	"net/http"
)

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

	result := string(body)
	if result == r {
		return "", true, nil
	}
	return result, false, nil
}
