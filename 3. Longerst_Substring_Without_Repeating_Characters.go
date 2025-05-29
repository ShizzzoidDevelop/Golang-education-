func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	buffer := len(s)

	for i := 0; i < buffer; i++ {
		seen := make(map[byte]bool)
		currentLength := 0

		for j := i; j < buffer; j++ {
			if seen[s[j]] {
				break
			}

			seen[s[j]] = true
			currentLength = j - i + 1

			if currentLength > maxLength {
				maxLength = currentLength
			}
		}
	}

	return maxLength
}