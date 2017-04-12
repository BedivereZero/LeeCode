/**
 * Return an array of size *returnSize.
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* twoSum(int* numbers, int numbersSize, int target, int* returnSize) {
	int* index = NULL;
	int i = 0;
	int j = 0;
	*returnSize = 2;
	index = (int *) malloc(*returnSize * sizeof(int));
	for (i=0; i < numbersSize - 1; i++) {
		for (j=i + 1; j < numbersSize; j++) {
			if (numbers[i] + numbers[j] == target) {
				index[0] = i + 1;
				index[1] = j + 1;
			}
		}
	}
	return index;
}
