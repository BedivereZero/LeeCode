class Solution:
    def mySqrt(self, x: int) -> int:
        if x < 2:
            return x
        l, r = 0, x
        while l < r - 1:
            # print('l:', l, 'r:', r)
            mid = (l + r) // 2
            # print('m:', mid)
            if mid * mid == x:
                return mid
            if mid * mid < x:
                l = mid
            else:
                r = mid
            # print('l:', l, 'r:', r)
        return l


if __name__ == "__main__":
    import math, sys
    for line in sys.stdin:
        print(int(Solution().mySqrt(int(line))))
    # for i in range(10):
    #     print(i, int(Solution().mySqrt(i)))
