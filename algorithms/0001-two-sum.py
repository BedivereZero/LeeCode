class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        results = dict()
        for index in range(len(nums)):
            if target - nums[index] in results:
                return [results[target - nums[index]], index]
            else:
                results[nums[index]] = index
