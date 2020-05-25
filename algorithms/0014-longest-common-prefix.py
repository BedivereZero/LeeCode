class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        if not strs:
            return ""
        length = len(strs[0])
        for each in strs[1:]:
            compare_len = min(length, len(each)
            for index in range(compare_len)):
                if strs[0][index] != each[index]:
                    length = index
                    break
            else:
                length = compare_len
        return strs[0][:length]
