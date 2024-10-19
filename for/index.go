package main

import "fmt"

func Map(strs []string, mapFunc func(s string) string) []string {
	res := []string{}

	for _, str := range strs {
		res = append(res, mapFunc(str))
	}

	return res
}

func exclamate(s string) string {
	return s + "!"
}

func main() {
	strs := []string{"Hello", "from", "go"}

	fmt.Print(strs)
	fmt.Print(Map(strs, exclamate))
}
