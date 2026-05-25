func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)
    var out []int
    for i,num := range nums {
        diff := target - num
        val, ok := seen[diff]
        if ok {
            out = append(out, val, i)
            return out
        } else {
            seen[num] = i
        }
    }
    return nil
}
