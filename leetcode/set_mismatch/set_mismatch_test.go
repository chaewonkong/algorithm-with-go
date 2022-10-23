package setmismatch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindErrorNums(t *testing.T) {
	assert.Equal(t, []int{2, 3}, FindErrorNums([]int{1, 2, 2, 4}))
	assert.Equal(t, []int{2, 4}, FindErrorNums([]int{1, 2, 2, 3, 5}))
}
