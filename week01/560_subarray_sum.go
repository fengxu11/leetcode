package week01

// 参考了官方的题解: 前缀和 + hash
func subarraySum(nums []int, k int) int {
	count, pre := 0, 0
	m := map[int]int{}
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := m[pre-k]; ok {
			count += m[pre-k]
		}
		// 记录前缀和出现的次数
		m[pre] += 1
	}
	return count
}
