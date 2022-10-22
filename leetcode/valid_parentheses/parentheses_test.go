package validparentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	assert.Equal(t, true, IsValid("()"))
	assert.Equal(t, true, IsValid("()[]"))
	assert.Equal(t, true, IsValid("({[]})"))
	assert.Equal(t, false, IsValid("(]"))
	assert.Equal(t, false, IsValid("({)}"))
	assert.Equal(t, false, IsValid("(){}}{"))
	assert.Equal(t, true, IsValid("(([]){})"))
}

func TestIsPair(t *testing.T) {
	p := newParentheses()
	assert.Equal(t, true, p.IsPair('(', ')'))
	assert.Equal(t, false, p.IsPair('(', '}'))
	assert.Equal(t, false, p.IsPair('(', '('))
	assert.Equal(t, false, p.IsPair(')', '('))
}

func TestIsOpening(t *testing.T) {
	p := newParentheses()
	assert.Equal(t, true, p.IsOpening('('))
	assert.Equal(t, true, p.IsOpening('{'))
	assert.Equal(t, false, p.IsOpening(')'))
	assert.Equal(t, false, p.IsOpening('}'))
}
