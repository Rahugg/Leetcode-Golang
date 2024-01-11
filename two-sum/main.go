package twosum

import "fmt"

func twoSum(nums []int, target int) []int {
	var (
		hashMap = map[int]int{}
		dif     int
	)
	for idx, val := range nums {
		dif = target - val
		idx2, ok := hashMap[dif]
		fmt.Println(dif, idx2, ok)
		if ok {
			return []int{idx, idx2}
		}
		hashMap[val] = idx
	}
	return []int{}
}