class Solution:
    def longestPalindrome(self, s: str) -> str:
        f = [[None for _ in range(len(s))]for _ in range(len(s))]
        a, b = 0, 0
        for length in range(len(s)):
            for head in range(len(s)):
                tail = head + length
                if tail < len(s):
                    if length < 2:
                        f[head][tail] = s[head] == s[tail]
                    else:
                        f[head][tail] = f[head + 1][tail - 1] and s[head] == s[tail]
                    if f[head][tail] and b - a < length:
                        a, b = head, tail
        return s[a:b + 1]
