// Time Complexity: O(n)
// Space Complexity: O(1)
// Runtime: 0 ms
// Memory: 10.23 MB

// Key Idea:
// 1. DP
// 2. Greedy
// 3. Kadane's Algorithm

func maxSubArray(nums []int) int {
	max_sum := nums[0]
	current_sum := nums[0]

	for _, num := range nums[1:] {
		current_sum = mymax(num, current_sum+num)
		max_sum = mymax(max_sum, current_sum)
	}

	return max_sum
}

func mymax(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

// Bad case (python)
// Time Complexity: O(n^3)

// nums = [-1, 3, 8, -4]
// max_len_array = len(nums)
// max_sum = float('-inf')

// for i in range(max_len_array):
//     target = i
//     while target < max_len_array:
//         total = 0
//         for j in range(i, target + 1):
//             total += nums[j]
//         if total > max_sum:
//             max_sum = total
//         target += 1
// print(max_sum)

// DP case (python)
// Time Complexity: O(n)

// nums = [-1, 3, 8, -4, 8, -9]
// dp = [0] * len(nums)
// dp[0] = nums[0]

// print(dp)

// for i in range(1, len(nums)):
//     dp[i] = max(nums[i], dp[i-1] + nums[i])
//     print(dp)

// max_sum = max(dp)
// print(max_sum)