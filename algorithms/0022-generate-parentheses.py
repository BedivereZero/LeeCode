class Solution:
    def search(self, result: List[int], current: str, left: int, right: int, remain: int):
        if not remain and left == right:
            result.append(current)
        if remain:
            self.search(result=result, current=current + '(', left=left + 1, right=right, remain= remain - 1)
        if right < left:
            self.search(result=result, current=current + ')', left=left, right=right + 1, remain=remain)

    def generateParenthesis(self, n: int) -> List[str]:
        result = list()
        self.search(result=result, current='', left=0, right=0, remain=n)
        return result
