void merge(int* nums1, int m, int* nums2, int n) {
	if (!(m && n))
		return;
	for (i=m; i > 1; i--)
		nums1[i] = nums1[i - 1];
	if (nums1[0] < nums2[0])
		nums1[1] = nums2[0];
	else {
		nums1[1] = nums1[0];
		nums1[0] = nums2[0];
	}
	return merge(nums1 + 2, m - 1, nums + 1, n - 1);
}
