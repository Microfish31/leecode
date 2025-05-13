// Time Complexity: O(n)
//   - Each number is processed at most once
// Space Complexity: O(n)
//   - A map is used to store all unique numbers
// Runtime: 59 ms
// Memory: 11.74 MB
//
// Key Idea:
// 1. Use a map (as a set) for O(1) lookups.
// 2. Only start counting a sequence from numbers that are the beginning of a sequence.
// 3. Avoid redundant work by ensuring each sequence is only explored once.
//
// Common Mistakes (That Cause O(n²) Time Complexity):
// 1. Scanning upward or downward from every number without skipping already processed parts
//    This results in repeatedly traversing the same sequences multiple times.
// 2. Forgetting to check if the number is the start of a sequence (`num - 1` not in the set)
//    Without this condition, the same sequence can be recomputed from different starting points.

func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	maxLen := 0

	for num := range numSet {
		if _, found := numSet[num-1]; !found {
			currentNum := num
			length := 1

			for {
				_, exists := numSet[currentNum+1]
				if !exists {
					break
				}
				currentNum++
				length++
			}

			if length > maxLen {
				maxLen = length
			}
		}
	}

	return maxLen
}
