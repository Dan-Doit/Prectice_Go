package main

import (
	"fmt"
	"strings"
)

func plus(a int, b int) int {
	return a + b
}

// export will start with Capital workds
func GetWords(words ...string) {
	// defer will run when this function is over
	defer plus(1, 2)
	fmt.Println(words)
}

// naked function
func lenUpper(word string) (length int, uppercase string) {
	length = len(word)
	uppercase = strings.ToUpper(word)
	return
}

func AddNum(numbers ...int) int {
	total := 0
	for index, number := range numbers {
		fmt.Println(index, number)
		total += number
	}
	return total
}

func main() {
	// var name string = "Dan"
	// func 안에서는 이런식으로 축약형을 쓸수있다.
	name := "Dan"
	name = "Jo"
	fmt.Println(name)
	fmt.Println(plus(2, 2))
	GetWords("Dan", "Bon", "ME", "Wine")
	fmt.Println(lenUpper("Good day"))
	fmt.Println(AddNum(1, 2, 3, 4, 5, 6, 7))
}
