package ispalindromeint

func IsPalindrome(x int) bool {
	// integers below 0 cannot be a palindrome
	if x < 0 {
		return false
	}

	// Find out total digits
	var ret = x
	var digits = 0
	for ret > 0 {
		ret /= 10
		digits++
	}

	s := make([]int, digits)

	// get numbers of digits
	for i := 0; i < digits; i++ {
		s[i] = x % 10
		x /= 10
	}

	// validation
	var i, j = 0, digits - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
