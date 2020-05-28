class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        nums_len = len(nums)
        if nums_len < 2:
            return nums_len
        p, q = 0, 1
        while q < nums_len:
            if nums[p] == nums[q]:
                q += 1
            else:
                p += 1
                nums[p] = nums[q]
                q += 1
        return p + 1
