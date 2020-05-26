class Solution:
    def isValid(self, s: str) -> bool:
        if not s:
            return True
        a = {
            '(': ')',
            '[': ']',
            '{': '}',
        }
        stack = list()
        for i in s:
            if i in a:
                stack.append(i)
                continue
            if not stack:
                return False
            if a[stack.pop()] != i:
                return False
        return not bool(stack)
