package easy_test

import (
	"my-leetcode/problems/easy"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuestion242(t *testing.T) {
	q := easy.NewQuestion242()

	t.Run("test 1", func(t *testing.T) {
		assert.True(t, q.IsAnagram("anagram", "nagaram"))
	})

	t.Run("test 2", func(t *testing.T) {
		assert.True(t, q.IsAnagram("aabb", "baba"))
	})
}
