// Time Complexity: O(mn)
// Space Complexity: O(1)
// Runtime: 0 ms
// Memory: 7.8 MB

// Key Idea:
// 1. Use the first row and column as markers for zeroes.
// 2. row0_flag and column0_flag record whether the first row and column need to be zeroed.
// 3. Do it in place

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	row0_flag := false    // True if the first row needs to be zeroed
	column0_flag := false // True if the first column needs to be zeroed

	// Check if the first column has any zero
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			column0_flag = true
			break
		}
	}

	// Check if the first row has any zero
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			row0_flag = true
			break
		}
	}

	// Use first row and column as markers for the rest of the matrix
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// Set zeroes based on markers
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// Zero out the first row if needed
	if row0_flag {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}

	// Zero out the first column if needed
	if column0_flag {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
