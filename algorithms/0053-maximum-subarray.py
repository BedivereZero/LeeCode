class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        length = len(nums)
        rec = nums[0]
        i = 1
        while i < length:
            if nums[i - 1] > 0:
                nums[i] = nums[i] + nums[i - 1]
            if rec < nums[i]:
                rec = nums[i]
            i += 1
        return rec
