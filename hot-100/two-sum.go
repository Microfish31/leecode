// Time Complexity: O(n) - Each element is visited once
// Space Complexity: O(n) - Hash map stores indices
// Runtime: 0 ms
// Memory: 5.92 MB

// Key Idea:
// 1. Use a hash map to store numbers and their indices.
// 2. Check if (target - current number) exists in the map.

func twoSum(nums []int, target int) []int {
	myMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		nowValue := nums[i]
		j, ok := myMap[target-nowValue]
		if ok {
			return []int{j, i}
		}
		myMap[nowValue] = i
	}
	return nil
}