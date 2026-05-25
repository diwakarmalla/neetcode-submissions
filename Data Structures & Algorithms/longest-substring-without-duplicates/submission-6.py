class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        count = 0
        for i in range(len(s)):
            substr = s[i]
            j = i+1
            while j < len(s):
                if s[j] not in substr:
                    substr += s[j]
                    j += 1
                else:
                    break
            count = max(count, len(substr))
            i = j
        return count
            
            



        