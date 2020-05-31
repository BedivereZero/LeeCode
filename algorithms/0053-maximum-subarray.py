class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        if len(nums) == 0:
            return 0
        if len(nums) == 1:
            return nums[0]
        mid = len(nums) // 2
        if nums[:mid]:
            lmax = nums[mid - 1]
            cur = nums[mid - 1]
            for i in range(mid - 2, -1, -1):
                cur += nums[i]
                if lmax < cur:
                    lmax = cur
        else:
            lmax = 0
        if nums[mid:]:
            rmax, cur = nums[mid], nums[mid]
            for i in range(mid + 1, len(nums)):
                cur += nums[i]
                if rmax < cur:
                    rmax = cur
        else:
            rmax = 0
        cmax = lmax + rmax
        return max(self.maxSubArray(nums[:mid]), cmax, self.maxSubArray(nums[mid:]))

if __name__ == "__main__":
    print(Solution().maxSubArray([1, 2]))
