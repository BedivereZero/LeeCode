class Solution:
    def addBinary(self, a: str, b: str) -> str:
        r = str()
        c = 0
        for i in range(max(len(a), len(b))):
            x = ord(a[len(a) - i - 1]) - ord('0') if i <= len(a) - 1 else 0
            y = ord(b[len(b) - i - 1]) - ord('0') if i <= len(b) - 1 else 0
            c, s = (x + y +c) // 2, (x + y + c) % 2
            r = chr(s + ord('0')) + r
        if c:
            r = '1' + r
        return r
