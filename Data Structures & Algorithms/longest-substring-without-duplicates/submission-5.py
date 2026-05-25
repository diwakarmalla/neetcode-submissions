class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        seen = ''
        count = 0
        for i in range(len(s)):
            seen += s[i]
            j = 1
            while i+j < len(s):
                if s[i+j] not in seen:
                    seen += s[i+j]
                    j += 1
                else:
                    seen=''
                    break
            count = max(count, j)
            i += j
        return count
            
            



        