bool containsNearbyDuplicate(int* nums, int numsSize, int k) {
	int i = 0;
	int j = 0;
	for (i=0; i < numsSize - 1; i++)
		for(j=i + 1; j < i + k; j++)
			if (nums[i] == nums[j])
				return true;
	return false;
}
