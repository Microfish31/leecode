// Time Complexity: O(n) - Each character is processed once
// Space Complexity: O(n) - Stack stores open brackets
// Runtime: 0 ms
// Memory: 4.62 MB

// Key Idea:
// 1. Use a stack to track open brackets.
// 2. Use a hash map for quick bracket pair lookup.

func isValid(s string) bool {
	stack := []string{}
	pairMap := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}

	for i := 0; i < len(s); i++ {
		str := s[i : i+1]
		if pair, ok := pairMap[str]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != pair {
				return false
			}
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, str)
	}

	return len(stack) == 0
}