class Solution:
    # def letterCombinations(self, digits: str) -> List[str]:
    def letterCombinations(self, digits):
        chs = {
            '2': 'abc',
            '3': 'def',
            '4': 'ghi',
            '5': 'jkl',
            '6': 'mno',
            '7': 'pqsr',
            '8': 'tuv',
            '9': 'wxyz',
        }
        out = list()
        for i in digits:
            cur = list()
            if i == '1':
                return []
            for c in chs[i]:
                if not out:
                    cur.append(c)
                for e in out:
                    cur.append(e + c)
            out = cur
        return out


if __name__ == "__main__":
    print(Solution().letterCombinations('123'))
