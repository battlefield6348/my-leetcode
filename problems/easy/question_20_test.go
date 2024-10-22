package easy_test

import (
	"my-leetcode/problems/easy"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuestion20(t *testing.T) {
	q := easy.NewQuestion20()
	t.Run("test 1", func(t *testing.T) {
		assert.True(t, q.IsValid("()"))
	})

	t.Run("test 2", func(t *testing.T) {
		assert.True(t, q.IsValid("()[]{}"))
	})

	t.Run("test 3", func(t *testing.T) {
		assert.False(t, q.IsValid("(]"))
	})

	t.Run("test 4", func(t *testing.T) {
		assert.True(t, q.IsValid("([])"))
	})

	t.Run("test 5", func(t *testing.T) {
		assert.False(t, q.IsValid("("))
	})

	t.Run("test 6", func(t *testing.T) {
		assert.False(t, q.IsValid("(("))
	})
}
