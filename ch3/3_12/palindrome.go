package palindrome

func IsPalindrome(s string) bool {
	beg := 0
	end := len(s) - 1

	for beg <= end {
		if s[beg] != s[end] {
			return false
		}

		beg += 1
		end -= 1
	}

	return true
}
