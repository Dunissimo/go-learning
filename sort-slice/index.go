package main

import (
	"fmt"
	"sort"
)

func UniqueSortedUserIDs(userIDs []int64) []int64 {
	res := []int64{}

	for _, n := range userIDs {
		if !Contains(n, res) {
			res = append(res, n)
		}
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})

	return res
}

func Contains(value int64, slice []int64) bool {
	for _, n := range slice {
		if n == value {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(UniqueSortedUserIDs([]int64{}))
	fmt.Println(UniqueSortedUserIDs([]int64{10}))
	fmt.Println(UniqueSortedUserIDs([]int64{55, 55}))
	fmt.Println(UniqueSortedUserIDs([]int64{55, 55, 33, 22}))
	fmt.Println(UniqueSortedUserIDs([]int64{55, 2, 88, 33, 2, 2, 55, 103, 33, 88}))
}
