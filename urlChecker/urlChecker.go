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
	status     = make(map[string]string)
)

func Checker() {
	state := "OK"
	status := status
	for _, url := range urls {
		err := getResponse(url)
		if err != nil {
			state = "FAILED"
		} else {
			state = "OK"
		}
		status[url] = state
	}
	fmt.Println("_______________________________________")
	for k, v := range status {
		fmt.Println(k, v)
	}
}

func getResponse(url string) error {
	fmt.Println("Checking : ", url)
	res, err := http.Get(url)
	if err != nil || res.StatusCode >= 400 {
		return errResFail
	} else {
		return nil
	}
}
