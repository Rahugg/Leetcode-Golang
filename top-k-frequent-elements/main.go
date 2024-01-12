package topkfrequentelements

import (
	"fmt"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)

	for _, val := range nums {
		m[val]++
	}
	// Create slice of key-value pairs
	pairs := make([][2]interface{}, 0, len(m))
	for k, v := range m {
		pairs = append(pairs, [2]interface{}{k, v})
	}

	// Sort slice based on values
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1].(int) > pairs[j][1].(int)
	})

	keys := make([]int, len(pairs))
	for i, p := range pairs {
		keys[i] = p[0].(int)
	}
	for _, k := range keys {
		fmt.Printf("%d: %d\n", k, m[k])
	}
	ans := make([]int, k)
	cnt := 0
	for _, val := range keys {
		if cnt == k {
			break
		}
		ans[cnt] = val
		cnt++
	}
	return ans
}
