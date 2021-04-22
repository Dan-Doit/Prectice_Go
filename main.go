package main

import (
	"fmt"

	"github.com/Dan-Doit/prectice-go/dicmap"
)

func main() {
	dics := dicmap.Dicts{}
	dics.AddDicts("Dan", "hi!")
	dics.AddDicts("Adam", "Good!")
	_, err := dics.UpdDicts("Con", "Good morning!")
	if err != nil {
		fmt.Println(err)
	}
	_, err = dics.UpdDicts("Adam", "Good morning!")
	if err != nil {
		fmt.Println(err)
	}

	dics.DelDicts("Dan")
	dics.Check("Dan")
}
