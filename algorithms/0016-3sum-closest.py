class Solution:
    def threeSumClosest(self, nums: List[int], target: int) -> int:
        nums.sort()
        closet_distance = None
        closet_target = None
        for a in range(len(nums)):
            if a > 0 and nums[a] == nums[a - 1]:
                continue
            b, c = a + 1, len(nums) - 1
            while b < c:
                if b > a + 1 and c < len(nums) - 1 and nums[b] == nums[b - 1] and nums[c] == nums[c + 1]:
                    b += 1
                    c -= 1
                    continue
                s = nums[a] + nums[b] + nums[c]
                distance = abs(s - target)
                if closet_distance is None or distance < closet_distance:
                    closet_distance, closet_target = distance, s
                if distance == 0:
                    return target
                if s < target:
                    b += 1
                else:
                    c -= 1
        return closet_target
