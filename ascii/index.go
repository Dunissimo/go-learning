package main

import (
	"fmt"
	"unicode"
)

func isASCII(s string) bool {
	for _, letter := range s {
		if letter > unicode.MaxASCII {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isASCII(" abc1"))
	fmt.Println(isASCII("хай"))
	fmt.Println(isASCII("hello \U0001F970"))
}
