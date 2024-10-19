package main

import "fmt"

func MergeNumberLists(numberLists ...[]int) []int {
	res := []int{}

	for _, v := range numberLists {
		res = append(res, v...)
	}

	return res
}

func main() {
	fmt.Print(MergeNumberLists([]int{1, 2}, []int{3}, []int{4}))
}
