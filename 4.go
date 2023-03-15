package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total_len := len(nums1) + len(nums2)
	if total_len%2 == 1 {
		k := total_len / 2
		return float64(getKthEle(nums1, nums2, k+1))
	} else {
		k := total_len / 2
		return float64(getKthEle(nums1, nums2, k+1)+getKthEle(nums1, nums2, k)) / 2.0
	}
}

func getKthEle(num1 []int, num2 []int, k int) int {
	idx1, idx2 := 0, 0
	for {
		if len(num1) == idx1 {
			return num2[idx2+k-1]
		}
		if len(num2) == idx2 {
			return num1[idx1+k-1]
		}
		if k == 1 {
			return min(num1[idx1], num2[idx2])
		}
		mid := k / 2
		newIdx1 := min(idx1+mid, len(num1)) - 1
		newIdx2 := min(idx2+mid, len(num2)) - 1
		pivot1, pivot2 := num1[newIdx1], num2[newIdx2]
		if pivot1 <= pivot2 {
			k = k - (newIdx1 - idx1 + 1)
			idx1 = newIdx1 + 1
		} else {
			k = k - (newIdx2 - idx2 + 1)
			idx2 = newIdx2 + 1
		}
	}
	return 0
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	num1 := []int{1, 3, 4, 5}
	num2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(findMedianSortedArrays(num1, num2))
}
