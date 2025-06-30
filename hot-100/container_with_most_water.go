// Runtime: 0 ms
// Memory: 9.77 MB
//
// Time Complexity: O(n)
//   - n: number of height bars
//   - Each bar is visited at most once by either the left or right pointer
//   - Total iterations are linear: O(n)
//
// Space Complexity: O(1)
//   - Only a constant amount of extra space is used for variables
//   - No additional data structures grow with input size
//
// Key Idea:
// 1. Use two pointers, starting from both ends of the array
// 2. Calculate the area between the two pointers based on the shorter height
// 3. Move the pointer with the shorter height inward to potentially find a larger area
// 4. Repeat until the two pointers meet, tracking the maximum area found

func maxArea(height []int) int {
	var index_low int
	var index_high int
	var max_area int
	index_low = 0
	index_high = len(height) - 1

	for index_low < index_high {
		height_index := 0
		distance := index_high - index_low

		if height[index_low] < height[index_high] {
			height_index = index_low
			index_low++
		} else {
			height_index = index_high
			index_high--
		}

		now_area := height[height_index] * distance
		if now_area > max_area {
			max_area = now_area
		}

		if index_low >= index_high {
			break
		}
	}

	return max_area
}
