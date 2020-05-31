class Solution:
    def plusOne(self, digits: List[int]) -> List[int]:
        digits.reverse()
        c = 1
        for i in range(len(digits)):
            digits[i], c = (digits[i] + c) % 10, (digits[i] + c) // 10
        if c:
            digits.append(1)
        digits.reverse()
        return digits
