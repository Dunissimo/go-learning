package main

import "fmt"

func MaxSum(nums1, nums2 []int) []int {
	sum1 := 0
	sum2 := 0

	for _, v := range nums1 {
		sum1 += v
	}

	for _, v := range nums2 {
		sum2 += v
	}

	if sum1 >= sum2 {
		return nums1
	} else {
		return nums2
	}
}

func main() {
	fmt.Println(MaxSum([]int{1, 2, 3}, []int{3, 4, 5, 6}))
	fmt.Println(MaxSum([]int{1, 2, 3}, []int{3, 2, 1}))
	fmt.Println(MaxSum([]int{10, 20}, []int{4, 6}))
}
