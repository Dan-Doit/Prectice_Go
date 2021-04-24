package urlChecker

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	urls = []string{
		"https://www.airbnb.com",
		"https://www.goolge.com",
		"https://www.naver.com",
		"https://www.kakao.com",
		"https://www.soundcloud.com",
		"https://www.orange.com",
		"https://www.amazon.com",
		"https://www.facebook.com",
		"https://www.fAsdasdasdasdbook.com",
	}
	errResFail = errors.New("Responce fail")
)

type requests struct {
	url    string
	status string
}

func Checker() {

	c := make(chan requests)
	responses := make(map[string]string)

	for _, v := range urls {
		go getResponse(v, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		responses[result.url] = result.status
	}

	for k, v := range responses {
		fmt.Println(k+" status : ", v)
	}
}

func getResponse(url string, c chan requests) {

	state := "OK"

	res, err := http.Get(url)
	if err != nil || res.StatusCode >= 400 {
		state = "FAIL"
	}

	c <- requests{url: url, status: state}
}
