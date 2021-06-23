package week01

// 合并两个数组
// 解题思路: i 和 j 两个索引指向的元素比较、谁小放谁、倒着放 节省一个数组的开辟、因为题目上是 num1后三个元素时空的
// 注意: 数组边界问题
func merge(num1 []int, m int, num2 []int, n int) {

	i := m - 1
	j := n - 1
	for k := m + n - 1; k >= 0; k-- {

		if j < 0 || (i >= 0 && num1[i] >= num2[j]) {
			num1[k] = num1[i]
			i--
		} else {
			num1[k] = num2[j]
			j--
		}
	}
}
