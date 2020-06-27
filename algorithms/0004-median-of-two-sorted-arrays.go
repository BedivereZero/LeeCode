package algorithms

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	a := find(nums1, nums2, (len(nums1)+len(nums2))/2)
	if (len(nums1)+len(nums2))%2 == 1 {
		return a
	}

	b := find(nums1, nums2, (len(nums1)+len(nums2))/2-1)
	return (a + b) / 2
}

func find(a []int, b []int, k int) float64 {
	if len(a) > len(b) {
		a, b = b, a
	}

	// one of a, b is empty
	if len(a) == 0 {
		return float64(b[k])
	}

	// k = 0
	if k == 0 {
		if a[0] < b[0] {
			return float64(a[0])
		} else {
			return float64(b[0])
		}
	}

	var ap, bp int
	if (k-1)/2 < len(a) {
		ap = (k - 1) / 2
	} else {
		ap = len(a) - 1
	}
	bp = (k - 1) / 2

	if a[ap] < b[bp] {
		return find(a[ap+1:], b, k-ap-1)
	} else {
		return find(a, b[bp+1:], k-bp-1)
	}
}
