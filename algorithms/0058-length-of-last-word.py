class Solution:
    def lengthOfLastWord(self, s: str) -> int:
        length = 0
        i = len(s) - 1
        while i >= 0 and s[i] == ' ':
            i -= 1
        end = i
        while i >= 0 and s[i] != ' ':
            i -= 1
        return end - i
