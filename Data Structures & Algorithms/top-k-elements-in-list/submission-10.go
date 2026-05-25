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
    countVal = countVal[len(countVal)-k:]
    for key, val := range count {
        fmt.Println(countVal)
        if slices.Contains(countVal, val) {
            out = append(out, key)
        }
    }
    return out
}
