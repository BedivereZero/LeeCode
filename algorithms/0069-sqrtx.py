class Solution:
    def mySqrt(self, x: int) -> int:
        if x < 2:
            return x
        a = (1 + x) / 2
        while True:
            b = 0.5 * (a + x / a)
            print('a:', a, 'b:', b)
            if int(a) == int(b):
                return int(a)
            a = b
