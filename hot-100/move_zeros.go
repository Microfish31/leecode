// Time Complexity: O(n)       // Only one pass through the array
// Space Complexity: O(1)      // In-place operation with constant extra space
// Runtime: 0 ms               // Based on LeetCode performance
// Memory: 8.88 MB              // Based on LeetCode performance

// Key Idea:
// 1. Use `izs` (index of zero start) to remember the first zero that needs to be swapped.
// 2. When a non-zero element is found, swap it with the zero at `izs` if `izs` is valid.
// 3. After a swap, move `izs` forward to find the next zero slot.
// 4. If `izs` no longer points to zero (i.e., was overwritten by a non-zero), reset it to -1.
// 5. This method preserves the relative order of non-zero elements and moves all zeroes to the end.

func moveZeroes(nums []int) {
	izs := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if izs == -1 {
				izs = i // Mark the first zero's position
			}
		} else {
			// Swap the non-zero element with the zero element found earlier
			if izs != -1 {
				nums[izs], nums[i] = nums[i], nums[izs]
				izs++ // Move izs forward to next potential zero
			}
		}

		// If the current izs no longer points to a zero, reset it
		if izs != -1 && nums[izs] != 0 {
			izs = -1
		}
	}
}

// method 2 (better readability)
// func moveZeroes(nums []int) {
//     insertPos := 0
//     for i := 0; i < len(nums); i++ {
//         if nums[i] != 0 {
//             nums[insertPos], nums[i] = nums[i], nums[insertPos]
//             insertPos++
//         }
//     }
// }

// method 3
// func moveZeroes(nums []int) {
// 	zeroCount := 0
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] == 0 {
// 			zeroCount++
// 		} else if zeroCount > 0 {
// 			nums[i-zeroCount], nums[i] = nums[i], nums[i-zeroCount]
// 		}
// 	}
// }
