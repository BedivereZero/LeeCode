package algorithms

func merge(nums1 []int, m int, nums2 []int, n int)  {
	nums3 := []int{}
	a, b := 0, 0
	for a < m && b < n {
		if nums1[a] < nums2[b] {
			nums3 = append(nums3, nums1[a])
			a++
		} else {
			nums3 = append(nums3, nums2[b])
			b++
		}
	}
	for a < m {
		nums3 = append(nums3, nums1[a])
		a++
	}
	for b < n {
		nums3 = append(nums3, nums2[b])
		b++
	}
	for i := 0; i < m + n; i++ {
		nums1[i] = nums3[i]
	}
}
