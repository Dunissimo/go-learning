package main

import "fmt"

func UniqueUserIDs(userIDs []int64) []int64 {
	res := []int64{}
	exist := make(map[int64]struct{})

	if len(userIDs) < 2 {
		return userIDs
	}

	for _, n := range userIDs {
		_, ok := exist[n]

		if ok {
			continue
		}

		exist[n] = struct{}{}
		res = append(res, n)
	}

	return res
}

func main() {
	fmt.Println(UniqueUserIDs([]int64{}))
	fmt.Println(UniqueUserIDs([]int64{10}))
	fmt.Println(UniqueUserIDs([]int64{55, 55}))
	fmt.Println(UniqueUserIDs([]int64{55, 55, 33, 22}))
	fmt.Println(UniqueUserIDs([]int64{55, 2, 88, 33, 2, 2, 55, 103, 33, 88}))
}
