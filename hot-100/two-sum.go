// Time complexity o(n)
// Space complexity o(n)
// Runtime 0 ms
// Memory 5.92 MB

// map store key-value
// key: input number value
// value: input number index

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