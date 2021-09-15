package snippets

func IsPalindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		// i-th element with last-i-1
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}
