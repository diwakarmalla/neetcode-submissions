class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        out = 1 if nums else 0
        for num in nums:
            i = 1
            while True:
                num1 = num+i
                if num1 in nums:
                    i += 1
                else:
                    break
                out = max(out, i)
        return out



        