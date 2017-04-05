class Solution {
	public:
		searchInsert(vector<int>& nums, int target) {
			int position = 0;
			while(position < nums.size() && nums[position] < target) {
				position++;
			}
			return position;
		}
};
