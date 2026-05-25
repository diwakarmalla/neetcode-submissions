class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        out = []
        nums.sort()
        for i in range(len(nums)-2):
            if i > 0 and nums[i] == nums[i-1]:
                continue
            for j in range(i+1, len(nums)-1):
                if j > i + 1 and nums[j] == nums[j-1]:
                    continue
                diff = 0-(nums[i]+nums[j])
                if diff in nums[j+1:]:
                    temp = [nums[i], nums[j], diff]
                    out.append(temp)
        return out
