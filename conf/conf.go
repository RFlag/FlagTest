package conf

import "net/http"

type Rule struct {
	Header http.Header
	Method string
	Url    string
	Result string
}

var Conf = map[string][]Rule{
	"login": {
		{
			Header: http.Header{
				"Authorization": []string{
					"JWT eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9." +
						"eyJleHAiOjE1MjUzNTIyMTQsInN1YiI6IiIsInVzZXJJZCI6MTE4MjksImFwcCI6IjEiLCJwbGEiOiJzaG9wIiwidmVyIjoiMSJ9." +
						"HC7pkiGgPQ_1PcbX1NnYnz9dOLgNVkLKo8lou6tLkcY=",
				},
			},
			Method: "POST",
			Url:    "/api/login?username=15137028515&password=1234567&app=1&platform=shop&version=1",
			Result: ``,
		},
	},
	"wallet": {
		{
			Header: http.Header{
				"Authorization": []string{
					"JWT eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjU4NTEyOTIsInN1YiI6IiIsInVzZXJJZCI6MTE4MjksImFwcCI6IjEiLCJwbGEi" +
						"OiJzaG9wIiwidmVyIjoiMSJ9._3X9YfDUzKuwsyU7JhYvsvB2msEueOq3WCzoEDBlt68=",
				},
			},
			Method: "POST",
			Url:    "/api/wallet/me",
			Result: `{"code":11,"error":"验证token失败","privateError":"token 已过期"}`,
		},
		{
			Header: http.Header{
				"Authorization": []string{
					"JWT eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9." +
						"eyJleHAiOjE1MjUzNTIyMTQsInN1YiI6IiIsInVzZXJJZCI6MTE4MjksImFwcCI6IjEiLCJwbGEiOiJzaG9wIiwidmVyIjoiMSJ9." +
						"HC7pkiGgPQ_1PcbX1NnYnz9dOLgNVkLKo8lou6tLkcY=",
				},
			},
			Method: "POST",
			Url:    "/api/wallet/history/list/all",
			Result: `{"code":11,"error":"验证token失败","privateError":"token 已过期"}`,
		},
	},
	"webapp": {
		{
			Header: http.Header{
				"TOKEN": []string{
					"JWT 2d782b8440adadf1801331572c7d8935",
				},
			},
			Method: "POST",
			Url:    "/session/login?mobile='15993929131'&password='123456'",
			Result: `{"code":1,"data":[{"Id":2,"CreateAt":"2018-04-28T20:58:52+08:00","DeleteAt":null,"App":"1","PrevId":1,` +
				`"UserId":11829,"OrderId":0,"ExtendId":0,"Extend":"","Name":"a","Money":100,"Amount":1,"Type":"aa","Memo":"bb"},` +
				`{"Id":7,"CreateAt":"2018-04-28T20:58:52+08:00","DeleteAt":null,"App":"1","PrevId":null,"UserId":11829,"OrderId":0,` +
				`"ExtendId":0,"Extend":"","Name":"interest","Money":1,"Amount":1,"Type":"付款","Memo":"付款1"}]}`,
		},
	},
}
