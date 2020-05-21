class Solution:
    def romanToInt(self, s: str) -> int:
        cha = {
            'I': 1,
            'V': 5,
            'X': 10,
            'L': 50,
            'C': 100,
            'D': 500,
            'M': 1000,
        }
        result = 0
        for i in range(len(s) - 1):
            if cha[s[i]] < cha[s[i + 1]]:
                result -= cha[s[i]]
            else:
                result += cha[s[i]]
        result += cha[s[-1]]
        return result
