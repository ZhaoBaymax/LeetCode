package mainfunc hash(x, y []int) map[int]int {	len_x := len(x)	res := map[int]int{}	for i := 0; i < len_x; i++ {		for j := 0; j < len_x; j++ {			res[x[i]+y[j]]++		}	}	return res}func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {	num_12 := hash(nums1, nums2)	ans := 0	for i := 0; i < len(nums1); i++ {		for j := 0; j < len(nums1); j++ {			tt := -nums3[i] - nums4[j]			ans += num_12[tt]		}	}	return ans}