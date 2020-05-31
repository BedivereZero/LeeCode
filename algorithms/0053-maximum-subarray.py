class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        if len(nums) == 0:
            return 0
        if len(nums) == 1:
            return nums[0]
        mid = len(nums) // 2
        lmax = max(sum(nums[i:mid]) for i in range(mid))
        rmax = max(sum(nums[mid:i]) for i in range(mid, len(nums) + 1))
        cmax = lmax + rmax
        return max(self.maxSubArray(nums[:mid]), cmax, self.maxSubArray(nums[mid:]))

if __name__ == "__main__":
    print(Solution().maxSubArray([1, 2]))
