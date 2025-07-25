// Time Complexity: O(n)
//   - n: number of characters in the string
//  - Each character is processed once
// 
// Space Complexity: O(Min(M,N))
//   - M: size of the character set (e.g., 256 for ASCII)
//   - N: length of the input string
// 
// Runtime: 0 ms
// Memory: 5 MB
// 
// Key Idea:
// 1. Use a map to store the last seen index of each character.
// 2. Maintain a sliding window [start, end] with no repeating characters.
// 3. If a character repeats inside the current window, move 'start' to (last seen index + 1).
// 4. Update the maximum length at each step based on the current window size.

func lengthOfLongestSubstring(s string) int {
    charIndex := make(map[byte]int)
	start := 0
	maxLen := 0

	for end := 0; end < len(s); end++ {
		char := s[end]
		if lastIndex, found := charIndex[char]; found && lastIndex >= start {
			start = lastIndex + 1
		}

		charIndex[char] = end

		if currentLen := end - start + 1; currentLen > maxLen {
			maxLen = currentLen
		}
	}
	return maxLen
}