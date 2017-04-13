int cmp(const int* a, const int* b) {
	return *a < *b ? -1 : 1;
}

bool containsDuplicate(int* nums, int numsSize) {
	int i = 0;
	qsort(nums, numsSize, sizeof(int), cmp);
	for (i=1; i < numsSize; i++) {
		if (nums[i] == nums[i - 1])
			return true;
	}
	return false;
}
