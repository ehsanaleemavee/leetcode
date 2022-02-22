class Solution:
    def titleToNumber(self, columnTitle: str) -> int:
        val = 0
        length = len(columnTitle)
        counter = length

        while counter > 0 :
            pos = length - counter
            ch = columnTitle[pos]
            val = val * 26 + ord(ch) -64
            counter = counter - 1
        
        return val