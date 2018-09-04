package main

import (
	"fmt"
	"sort"
)

func arrayPairSum(nums []int) int {
	var sum = 0
	sort.Ints(nums)
	for i := 0; i < len(nums); i+=2 {
		sum += nums[i]
	}
	return sum
}

func main() {
	n := []int {
		-100, -1, -2, -3, 0, 1, 1, 4, 3, 2,
	}
	fmt.Println(arrayPairSum(n))
}
