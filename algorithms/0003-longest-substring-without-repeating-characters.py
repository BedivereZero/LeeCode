class Solution:

    def lengthOfLongestSubstring(self, s: str) -> int:

        l, x, y = 0, 0, 0

        while True:

            # get longest substing without repeating characters startswith index x
            while True:
                y += 1
                if len(s) < y:
                    l = max(l, y - x - 1)
                    return l
                if len(set(s[x:y])) != len(s[x:y]):
                    break
            l = max(l, y - x - 1)

            # get longest substing without repeating characters endswith index y
            while True:
                x += 1
                if s[x - 1] == s[y - 1]:
                    break
