class Solution:
    def isValid(self, s: str) -> bool:
        chars = {')':'(', '}':'{', ']':'['}
        stack = []
        for c in s:
            if stack and c in chars:
                temp = stack.pop()
                if chars[c] != temp:
                    return False
            else:
                stack.append(c)
        return not stack
                

        