package main

import "fmt"

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int, len(nums))

	for i, v := range nums {
		mp[v] = i
	}

	var candidate int
	for i, v := range nums {
		candidate = target - v
		if idx, ok := mp[candidate]; ok && i != idx {
			return []int{i, idx}
		}
	}

	return []int{-1, -1}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 18
	fmt.Println(twoSum(nums, target))
}
