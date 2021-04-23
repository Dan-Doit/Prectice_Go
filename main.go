package main

import (
	"fmt"
	"time"

	"github.com/Dan-Doit/prectice-go/urlChecker"
)

func main() {
	go gorutine("first")
	go gorutine("second")
	gorutine("third")
	urlChecker.Checker()

}

func gorutine(word string) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println(word, i)
	}
}
