// Time Complexity: O(n * k)
//   - n: number of input strings
//   - k: average length of each string
//   - Counting characters takes O(k), total O(n * k)
// Space Complexity: O(n * k)
//   - For storing groups of anagrams in a map
//   - Map keys use fixed size (26 bytes per key), values store original strings

// Key Idea:
// 1. For each word, count the frequency of each character (a-z) in a 26-byte array
// 2. Use the byte array (converted to string) as a unique key to group anagrams
// 3. This avoids sorting and gives better performance than O(k log k) approaches

func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string)

	for _, str := range strs {
		count := [26]byte{}
		for i := 0; i < len(str); i++ {
			count[str[i]-'a']++
		}

		key := string(count[:])

		anagramMap[key] = append(anagramMap[key], str)
	}

	result := make([][]string, 0, len(anagramMap))
	for _, group := range anagramMap {
		result = append(result, group)
	}

	return result
}
