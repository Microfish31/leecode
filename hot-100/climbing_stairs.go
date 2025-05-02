// Time Complexity: O(n)
// Space Complexity: O(1)
// Runtime: 0 ms
// Memory: 3.78 MB

// Key Idea:
// 1. Enumerate all combinations of 1-step and 2-step
// 2. Use combination formula to count permutation (better than factorial)

func climbStairs(n int) int {
	sum := 0
	for twoSteps := 0; twoSteps <= n/2; twoSteps++ {
		oneSteps := n - 2*twoSteps
		totalSteps := oneSteps + twoSteps        //(1 counts,2 counts)
		sum += combination(totalSteps, twoSteps) // c(all counts,2 )
	}
	return sum
}

func combination(n, k int) int {
	if k == 0 || k == n {
		return 1
	}
	res := 1
	for i := 1; i <= k; i++ {
		res = res * (n - i + 1) / i
	}
	return res
}

