class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        ln = len(nums)
        nums.sort()
        r = list()
        a = 0
        while a < ln:
            if a > 0 and nums[a] == nums[a - 1]:
                a += 1
                continue
            b, c = a + 1, ln - 1
            while b < c:
                if b > a + 1 and c < ln - 1 and nums[b] == nums[b - 1] and nums[c] == nums[c + 1]:
                    b += 1
                    c -= 1
                    continue
                if nums[a] + nums[b] + nums[c] < 0:
                    b += 1
                elif nums[a] + nums[b] + nums[c] > 0:
                    c -= 1
                else:
                    r.append([nums[a], nums[b], nums[c]])
                    b += 1
                    c -= 1
            a += 1
        return r
