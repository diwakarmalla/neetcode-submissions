import (
    "maps"
    "slices"
)

func topKFrequent(nums []int, k int) []int {
    count := make(map[int]int)
    out := []int{}

    for _, num := range nums {
        count[num]++
    }

    countVal := slices.Collect(maps.Values(count))
    slices.Sort(countVal)

    // take largest k frequencies
    countVal = countVal[len(countVal)-k:]

    // convert slice to set for lookup
    freqSet := make(map[int]bool)
    for _, v := range countVal {
        freqSet[v] = true
    }

    for key, val := range count {
        if freqSet[val] {
            out = append(out, key)
        }
    }

    return out
}