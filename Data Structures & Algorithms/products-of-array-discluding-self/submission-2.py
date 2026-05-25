class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        out = []
        for i, num in enumerate(nums):
            out.append(math.prod(nums[:i]+nums[i+1:]))
        return out


        