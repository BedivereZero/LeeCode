class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        if not nums:
            return 0
        l, r = 0, len(nums) - 1
        while l < len(nums) and nums[l] != val:
            l += 1
        while 0 <= r and nums[r] == val:
            r -= 1
        while l < r:
            nums[l], nums[r] = nums[r], nums[l]
            while nums[l] != val:
                l += 1
            while nums[r] == val:
                r -= 1
        return r + 1
