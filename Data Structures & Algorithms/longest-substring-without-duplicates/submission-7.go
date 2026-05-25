import "slices"
func lengthOfLongestSubstring(s string) int {
    count := 0
    for i, char := range s {
        substr := []rune{rune(char)}
        j := i+1
        for j < len(s) {
            if !slices.Contains(substr, rune(s[j])) {
                substr = append(substr, rune(s[j]))
                j += 1
            } else {
                break
            }

        }
        i = j
        count = max(count, len(substr))
    }
    return count

}
