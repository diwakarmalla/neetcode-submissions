func isValid(s string) bool {
    stack := []rune{}

    closeToOpen := map[rune]rune{
        ')': '(',
        ']': '[',
        '}': '{',
    }

    for _, c := range s {
        open, exists := closeToOpen[c]

        if exists {
            // stack empty
            if len(stack) == 0 {
                return false
            }

            // pop
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]

            if top != open {
                return false
            }
        } else {
            // push
            stack = append(stack, c)
        }
    }

    return len(stack) == 0
}