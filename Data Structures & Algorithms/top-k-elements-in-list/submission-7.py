class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        count=defaultdict(int)
        out=[]
        for num in nums:
            count[num] += 1
        temp = sorted(count.values(), reverse=True)[:k]
        for key, val in count.items():
            if val in temp:
                out.append(key)
        return out


        