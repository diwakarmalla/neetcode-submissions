class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        anagrams=defaultdict(list)
        for temp in strs:
            key = ''.join(sorted(temp))
            anagrams[key].append(temp)
        return list(anagrams.values())

        