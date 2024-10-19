package main

import (
	"fmt"
	"time"
)

func MaxSum(nums1, nums2 []int) []int {
	s1, s2 := 0, 0

	go SumParallel(nums1, &s1)
	go SumParallel(nums2, &s2)

	time.Sleep(100 * time.Millisecond)

	if s1 >= s2 {
		return nums1
	}

	return nums2
}

func SumParallel(nums []int, res *int) {
	*res = Sum(nums)
}

func Sum(nums []int) int {
	s := 0

	for _, v := range nums {
		s += v
	}

	return s
}

func main() {
	fmt.Println(MaxSum([]int{1, 2, 3}, []int{3, 4, 5, 6}))
	fmt.Println(MaxSum([]int{1, 2, 3}, []int{3, 2, 1}))
	fmt.Println(MaxSum([]int{10, 20}, []int{4, 6}))

	time.Sleep(100 * time.Millisecond)
}
