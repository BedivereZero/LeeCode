class Solution:
    def maxArea(self, height: List[int]) -> int:
        a, l, r = 0, 0, len(height) - 1
        while l < r:
            h = min(height[l], height[r])
            ca = h * (r - l)
            if a < ca:
                a = ca
            while l < len(height) and height[l] <= h:
                l += 1
            while r >= 0 and height[r] <= h:
                r -= 1
        return a
