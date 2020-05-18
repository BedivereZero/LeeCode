class Solution:
    def reverse(self, x: int) -> int:
        if x < - (1 << 31) or not x < (1 << 31) - 1:
            return 0
        negative = x < 0
        strx = str(-x) if negative else str(x)
        stry = strx[::-1]
        if negative:
            stry = '-' + stry
        y = int(stry)
        if y < - (1 << 31) or not y < (1 << 31) - 1:
            y = 0
        return y