package easy_test

import (
	"my-leetcode/problems/easy"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuestion217(t *testing.T) {
	q := easy.NewQuestion217()

	t.Run("test 1", func(t *testing.T) {
		assert.True(t, q.ContainsDuplicate([]int{1, 2, 3, 1}))
	})

	t.Run("test 2", func(t *testing.T) {
		assert.False(t, q.ContainsDuplicate([]int{1, 2, 3, 4}))
	})
}
