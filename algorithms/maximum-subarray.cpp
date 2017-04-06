class Solution {
public:
    int maxSubArray(vector<int>& nums) {
		int **a = new int*[nums.size() - 1];
		for (int i=0; i < nums.size() - 1; i++)
			a[i] = new int[nums.size() - 1];
		for (int p=0; p < nums.size() - 1; p++) {
			for (int l=1; l < nums.size() - p; l++) {
				a[p][l] = 0;
				for(int i=p; i < p + l; i++)
					a[p][l] += nums[i];
			}
		}
		int max;
		bool found = false;
		for (int p=0; p < nums.size() - 1; p++) {
			for (int l=1; l < nums.size() - p; l++) {
				if (!found) {
					found = true;
					max = a[p][l];
				}
				else {
					if (max < a[p][l])
						max = a[p][l];
				}
			}
		}
		return max;
	}
};
