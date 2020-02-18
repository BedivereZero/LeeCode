class Solution:

    def lengthOfLongestSubstring(self, s: str) -> int:

        l, x, y = 0, 0, 0
        m = set()

        while True:

            # get longest substing without repeating characters startswith index x
            while True:
                y += 1
                if len(s) < y:
                    l = max(l, y - x - 1)
                    # print('C: %3d, %3d, %3d, %s%s%s, %s' % (l, x, y, '_' * x, s[x:y], '_' * (len(s) - y), m))
                    return l
                if s[y - 1] in m:
                    break
                else:
                    m.add(s[y - 1])
            l = max(l, y - x - 1)
            # print('A: %3d, %3d, %3d, %s%s%s, %s' % (l, x, y, '_' * x, s[x:y], '_' * (len(s) - y), m))

            # get longest substing without repeating characters endswith index y
            while True:
                x += 1
                if s[x - 1] == s[y - 1]:
                    break
                m.remove(s[x - 1])
            # print('B: %3d, %3d, %3d, %s%s%s, %s' % (l, x, y, '_' * x, s[x:y], '_' * (len(s) - y), m))


if __name__ == '__main__':
    Solution().lengthOfLongestSubstring('abcabcbb')
