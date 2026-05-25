import "slices"
func longestConsecutive(nums []int) int {
    count := 0
    for _, num := range nums {
        if !slices.Contains(nums, num-1) {
            length := 1
            for {
                if slices.Contains(nums, num+length) {
                    length += 1
                } else {
                    break
                }

            }
            count = max(count, length)

        }
    }
    return count
    

}
