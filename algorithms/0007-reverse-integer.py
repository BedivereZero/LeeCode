class Solution:
    def reverse(self, x: int) -> int:
        negative = x < 0
        if negative:
            x = -x
        if x > 5927694924:
            return 0
        y = 0
        while x:
            y = y * 10 + x % 10
            x = x // 10
        y = -y if negative else y
        if y > (1 << 31) - 1:
            return 0
        if y < - (1 << 31):
            return 0
        return y