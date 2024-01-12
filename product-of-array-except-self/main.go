package productofarrayexceptself

// TC-  O(N) MC - O(n)
func productExceptSelf(nums []int) []int {
	pre := make([]int, len(nums))
	suf := make([]int, len(nums))
	ans := make([]int, len(nums))

	pre[0] = 1
	suf[len(nums)-1] = 1

	for i := 1; i < len(nums); i++ {
		pre[i] = pre[i-1] * nums[i-1]
	}

	for i := len(nums) - 2; i >= 0; i-- {
		suf[i] = suf[i+1] * nums[i+1]
	}

	for i := 0; i < len(nums); i++ {
		ans[i] = pre[i] * suf[i]
	}

	return ans
}

// TC - O(N) MC - O(1)
// func productExceptSelf(nums []int) []int {
// 	n := len(nums)
// 	ans := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		ans[i] = 1
// 	}
// 	curr := 1
// 	for i := 0; i < n; i++ {
// 		ans[i] *= curr
// 		curr *= nums[i]
// 	}
// 	curr = 1
// 	for i := n - 1; i >= 0; i-- {
// 		ans[i] *= curr
// 		curr *= nums[i]
// 	}
// 	return ans
// }
