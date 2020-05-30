class Solution:
    def countAndSay(self, n: int) -> str:
        say = '1'
        for _ in range(1, n):
            cur = str()
            char = None
            count = 0
            for c in say:
                if char == c:
                    count += 1
                else:
                    if count:
                        cur += str(count) + char
                    char = c
                    count = 1
            if count:
                cur += str(count) + char
            say = cur
        return say


if __name__ == "__main__":
    print(Solution().countAndSay(n=2))
