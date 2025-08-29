// Time Complexity: o(log(n))
// Space Complexity: o(1)
// Runtime: 0 ms
// Memory: 4.82 MB

// Key Idea:
// 1. binary search
// 2. left, right, mid variables
// 3. Don't need to use tree structure (worse cache performance)

func searchInsert(nums []int, target int) int {
    left,right := 0, len(nums)-1

    for left <= right {
        mid := (left + right) / 2

        // find
        if nums[mid] == target {
            return mid
        // smaller
        }else if nums[mid] < target {
            left = mid + 1
        // bigger
        }else {
            right = mid - 1
        }
    }

	// not find
    return left
}