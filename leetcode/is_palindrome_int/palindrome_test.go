package ispalindromeint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	// 1 digit numbers are always a palindrome
	assert.Equal(t, true, IsPalindrome(3))
	assert.Equal(t, true, IsPalindrome(0))

	// 2 digit number that not a palindrome
	assert.Equal(t, false, IsPalindrome(10))

	// 2 digit number that is palindrome
	assert.Equal(t, true, IsPalindrome(11))

	// 3 digit number that is palindrome
	assert.Equal(t, true, IsPalindrome(121))

	// 3 digit number that not a palindrome
	assert.Equal(t, false, IsPalindrome(123))

	// negative integers never gonna be  palindromes
	assert.Equal(t, false, IsPalindrome(-121))
}
