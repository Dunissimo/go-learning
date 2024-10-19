package main

import (
	"fmt"
	"strings"
	"unicode"
)

func latinLetters(s string) string {
	sb := &strings.Builder{}

	for _, letter := range s {
		if unicode.Is(unicode.Latin, letter) {
			sb.WriteString(string(letter))
		}
	}

	return sb.String()
}

func main() {
	fmt.Println(latinLetters(latinLetters(" abc1")))
	fmt.Println(latinLetters(latinLetters(" привет")))
	fmt.Println(latinLetters(latinLetters("11 ! t e    s t %)")))
	fmt.Println(latinLetters(latinLetters("John Уолтер")))
	fmt.Println(latinLetters(latinLetters("wo[r]d")))
}
