package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	str := strings.Join(os.Args[1:], " ")
	if str == "" {
		fmt.Println(0)
		return
	}
	wordCount := strings.Split(str, " ")
	fmt.Println(len(wordCount))
}
