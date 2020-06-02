package algorithms

func merge(nums1 []int, m int, nums2 []int, n int)  {
	a, b := m - 1, n - 1
	for a >= 0 && b >= 0 {
		if nums1[a] > nums2[b] {
			nums1[a + b + 1] = nums1[a]
			a--
		} else {
			nums1[a + b + 1] = nums2[b]
			b--
		}
	}
	for b >= 0 {
		nums1[b] = nums2[b]
		b--
	}
}
