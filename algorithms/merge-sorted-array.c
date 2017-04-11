void merge(int* nums1, int m, int* nums2, int n) {
	int i = 0;
	int j = 0;
	if (!n)
		return;
	while (i < m) {
		if (nums2[0] < nums1[i])
			break;
		else
			i++;
	}
	for (j=m; j > i; j--)
		nums1[j] = nums1[j - 1];
	nums1[i] = nums2[0];
	return merge(nums1 + i + 1, m - i, nums2 + 1, n - 1);
}
