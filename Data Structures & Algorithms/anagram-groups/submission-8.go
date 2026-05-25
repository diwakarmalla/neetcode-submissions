import (
    "maps"
    "slices"
)
func groupAnagrams(strs []string) [][]string {
    anagrams := make(map[string][]string)
    for _, temp := range strs {
        r := []rune(temp)
        slices.Sort(r)
        anagrams[string(r)] = append(anagrams[string(r)], temp)
    }
    return slices.Collect(maps.Values(anagrams))

    

}
