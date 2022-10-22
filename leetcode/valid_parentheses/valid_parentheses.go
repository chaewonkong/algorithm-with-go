package validparentheses

// https://leetcode.com/problems/valid-parentheses/submissions/

type Parentheses struct {
	parenMap map[rune]rune
}

func newParentheses() Parentheses {
	return Parentheses{
		parenMap: map[rune]rune{
			'(': ')',
			'{': '}',
			'[': ']',
		},
	}
}

func IsValid(s string) bool {
	p := newParentheses()
	stack := make([]rune, 0)

	for _, char := range s {
		if p.IsOpening(char) {
			stack = append(stack, char)
			// list.PushBack(char)
		} else if len(stack) > 0 {
			last := stack[len(stack)-1]
			if p.IsPair(last, char) {
				// stack.pop()
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			return false
		}
	}
	if len(stack) > 0 {
		return false
	}

	return true
}

func (p Parentheses) IsPair(s1 rune, s2 rune) bool {
	return p.parenMap[s1] == s2
}

func (p Parentheses) IsOpening(s rune) bool {
	if _, ok := p.parenMap[s]; ok {
		return true
	}
	return false
}
