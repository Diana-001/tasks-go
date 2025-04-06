package main

import (
	"fmt"
	"tasks-go/words"
)

func main() {
	a := words.FindAnagrams()

	fmt.Println(a)
}
