class Solution:
    def myAtoi(self, str: str) -> int:
        max_int, min_int = 2147483647, -2147483648
        for i in range(len(str)):
            if str[i] != ' ':
                h = i
                break
        else:
            return 0

        if str[h] not in '+-0123456789':
            return 0

        negative = str[h] == '-'
        if str[h] in '+-':
            h = h + 1

        for i in range(h, len(str)):
            if str[i] not in '0123456789':
                t = i
                break
        else:
            t = len(str)
        if h == t:
            return 0
        print('h:', h, 't:', t)
        integer = -1 * (int(negative) * 2 - 1) * (ord(str[h]) - ord('0'))
        print('integer:', integer)
        for c in str[h+1:t]:
            i = ord(c) - ord('0')
            print('c:', c, 'i:', i)
            if negative:
                if integer < min_int / 10:
                    return min_int
                else:
                    integer = integer * 10
                if integer < min_int + i:
                    return min_int
                else:
                    integer = integer - i
            else:
                if integer > max_int / 10:
                    return max_int
                else:
                    integer = integer * 10
                if integer > max_int - i:
                    return max_int
                else:
                    integer = integer + i
        return integer
