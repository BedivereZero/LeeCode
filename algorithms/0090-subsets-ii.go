package algorithms

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	subsets := [][]int{{}}
	a := 0
	for a < len(nums) {
		b := a
		var cur [][]int
		for b < len(nums) && nums[a] == nums[b] {
			var src [][]int
			// 当前数字的解集
			if len(cur) == 0 {
				src = subsets
			} else {
				src = cur
			}
			cur = make([][]int, len(src))
			for i := range cur {
				cur[i] = make([]int, len(src[i])+1)
				copy(cur[i], src[i])
				cur[i][len(cur[i])-1] = nums[a]
			}
			// 追加至全部数字的解集
			subsets = append(subsets, cur...)
			b++
		}
		a = b
	}
	return subsets
}
