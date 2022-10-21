package main

//给你一个大小为 n 下标从 0 开始的整数数组 nums 和一个正整数 k 。
//
//对于 k <= i < n - k 之间的一个下标 i ，如果它满足以下条件，我们就称它为一个 好 下标：
//
//下标 i 之前 的 k 个元素是 非递增的 。
//下标 i 之后 的 k 个元素是 非递减的 。
//按 升序 返回所有好下标。
//
//
//
//示例 1：
//
//输入：nums = [2,1,1,1,3,4,1], k = 2
//输出：[2,3]
//解释：数组中有两个好下标：
//- 下标 2 。子数组 [2,1] 是非递增的，子数组 [1,3] 是非递减的。
//- 下标 3 。子数组 [1,1] 是非递增的，子数组 [3,4] 是非递减的。
//注意，下标 4 不是好下标，因为 [4,1] 不是非递减的。
//示例 2：
//
//输入：nums = [2,1,1,2], k = 2
//输出：[]
//解释：数组中没有好下标。

func goodIndices(nums []int, k int) []int {
	len_num := len(nums)
	forward := make([]int, len_num)
	backward := make([]int, len_num)
	var res []int
	k = k - 1
	for i := 1; i < len_num; i++ {
		if nums[i] <= nums[i-1] {
			forward[i] = forward[i-1] + 1
		}
	}
	for i := len_num - 2; i >= 0; i-- {
		if nums[i] <= nums[i+1] {
			backward[i] = backward[i+1] + 1
		}
	}
	for i := k + 1; i < len_num-k-1; i++ {
		if forward[i-1] >= k && backward[i+1] >= k {
			res = append(res, i)
		}
	}
	return res
}

func main() {
	s := []int{2, 1, 1, 1, 3, 4, 1}
	goodIndices(s, 2)
}
