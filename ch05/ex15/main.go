package main

import "fmt"

func max(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}

	res := nums[0]
	for _, v := range nums {
		if v > res {
			res = v
		}
	}
	return res
}

func max2(essential int, nums ...int) int {
	res := essential
	for _, v := range nums {
		if v > res {
			res = v
		}
	}
	return res
}

func min(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}

	res := nums[0]
	for _, v := range nums {
		if v < res {
			res = v
		}
	}
	return res
}

func min2(essential int, nums ...int) int {
	res := essential
	for _, v := range nums {
		if v < res {
			res = v
		}
	}
	return res
}

func main() {
	fmt.Println(max(-2, 4, -8, 16))
	fmt.Println(min(-2, 4, -8, 16))
	fmt.Println(max2(-2, 4, -8, 16))
	fmt.Println(min2(-2, 4, -8, 16))
}
