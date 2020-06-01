class Solution:
    def reverse_str(self, input: str) -> str:
        if input:
            return self.reverse_str(input[1:]) + input[0]
        else:
            return input
    
    def addBinary(self, a: str, b: str) -> str:
        a = self.reverse_str(a)
        b = self.reverse_str(b)
        t = {
            ('0', '0', '0'): ('0', '0'),
            ('0', '0', '1'): ('0', '1'),
            ('0', '1', '0'): ('0', '1'),
            ('0', '1', '1'): ('1', '0'),
            ('1', '0', '0'): ('0', '1'),
            ('1', '0', '1'): ('1', '0'),
            ('1', '1', '0'): ('1', '0'),
            ('1', '1', '1'): ('1', '1'),
        }
        r = str()
        i = 0
        c = '0'
        while i < len(a) or i < len(b):
            x = a[i] if i < len(a) else '0'
            y = b[i] if i < len(b) else '0'
            c, s = t[(x, y, c)]
            r = r + s
            i += 1
        if c == '1':
            r = r + c
        return self.reverse_str(r)
