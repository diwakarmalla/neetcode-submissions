func productExceptSelf(nums []int) []int {
    out := []int{}
    for i, _ := range nums {
        product := 1
        newSlice := make([]int, len(nums))
        copy(newSlice, nums)
        newSlice = append(newSlice[:i], newSlice[i+1:]...)
        for _, temp := range newSlice {
            product *= temp
        }
        out = append(out, product)
    }
    return out

}
