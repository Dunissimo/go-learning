package main

import (
	"fmt"
)

func IntsCopy(src []int, maxLen int) []int {
	if maxLen <= 0 {
		return []int{}
	}

	var res []int

	if maxLen > len(src) {
		res = make([]int, len(src))

	} else {
		res = make([]int, maxLen)
	}

	copy(res, src)

	return res
}

func main() {
	fmt.Println(IntsCopy([]int{}, 0))
	fmt.Println(IntsCopy([]int{1, 2}, 0))
	fmt.Println(IntsCopy([]int{1, 2}, -1))
	fmt.Println(IntsCopy([]int{1, 2}, -5))
	fmt.Println(IntsCopy([]int{1, 2, 3}, 2))
	fmt.Println(IntsCopy([]int{1, 2, 3, 4}, 4))
	fmt.Println(IntsCopy([]int{1, 2, 3, 4, 5}, 10))
}
