class Solution:
    def fourSum(self, nums: List[int], target: int) -> List[List[int]]:
        nums.sort()
        r = list()
        for a in range(len(nums)):
            if a > 0 and nums[a] == nums[a - 1]:
                continue
            for b in range(a + 1, len(nums)):
                if b > a + 1 and nums[b] == nums[b - 1]:
                    continue
                c, d = b + 1, len(nums) - 1
                while c < d:
                    if c > b + 1 and d < len(nums) - 1 and nums[c] == nums[c - 1] and nums[d] == nums[d + 1]:
                        c += 1
                        d -= 1
                        continue
                    if nums[a] + nums[b] + nums[c] + nums[d] < target:
                        c += 1
                    elif nums[a] + nums[b] + nums[c] + nums[d] > target:
                        d -= 1
                    else:
                        r.append([nums[a], nums[b], nums[c], nums[d]])
                        c += 1
                        d -= 1
        return r
