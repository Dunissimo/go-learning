package main

import "fmt"

func shiftASCII(s string, step int) string {
	if step == 0 {
		return s
	}

	res := make([]byte, len(s))

	for i, letter := range s {
		res[i] = byte(letter) + byte(step)
	}

	return string(res)
}

func main() {
	fmt.Println(shiftASCII("abc", 0))
	fmt.Println(shiftASCII("abc1", 1))
	fmt.Println(shiftASCII("bcd2", -1))
	fmt.Println(shiftASCII("hi", 10))
	fmt.Println(shiftASCII("abc", 256))
	fmt.Println(shiftASCII("abc", -511))
}
