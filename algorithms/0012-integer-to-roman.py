class Solution:
    def intToRoman(self, num: int) -> str:
        r = str()

        c, num = num // 1000, num % 1000
        r += 'M' * c

        c, num = num // 100, num % 100
        if c == 9:
            r += 'CM'
        elif c in range(5, 9):
            r += 'D' + 'C' * (c - 5)
        elif c == 4:
            r += 'CD'
        else:
            r += 'C' * c

        c, num = num // 10, num % 10
        if c == 9:
            r += 'XC'
        elif c in range(5, 9):
            r += 'L' + 'X' * (c - 5)
        elif c == 4:
            r += 'XL'
        else:
            r += 'X' * c

        c = num
        if c == 9:
            r += 'IX'
        elif c in range(5, 9):
            r += 'V' + 'I' * (c - 5)
        elif c == 4:
            r += 'IV'
        else:
            r += 'I' * c

        return r
