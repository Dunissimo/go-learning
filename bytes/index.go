package main

import "fmt"

func nextASCII(b byte) byte {
	return b + 1
}

func prevASCII(b byte) byte {
	return b - 1
}

func main() {
	fmt.Println(string(nextASCII(byte('b'))))

	fmt.Println(string(prevASCII(byte('c'))))
}
