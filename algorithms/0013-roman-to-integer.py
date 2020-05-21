class Solution:
    def romanToInt(self, s: str) -> int:
        char = {
            'I': 1,
            'V': 5,
            'X': 10,
            'L': 50,
            'C': 100,
            'D': 500,
            'M': 1000,
            'IV': 4,
            'IX': 9,
            'XL': 40,
            'XC': 90,
            'CD': 400,
            'CM': 900,
        }
        length = len(s)
        integer = 0
        p = 0
        while p < length:
            if (length - p) < 2:
                integer += char[s[p]]
                break
            if s[p:p+2] in char:
                integer += char[s[p:p+2]]
                p += 2
            else:
                integer += char[s[p]]
                p += 1
        return integer
