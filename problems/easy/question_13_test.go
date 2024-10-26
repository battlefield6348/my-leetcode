package easy_test

import (
	"my-leetcode/problems/easy"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuestion13(t *testing.T) {
	q := easy.NewQuestion13()
	t.Run("test 1", func(t *testing.T) {
		assert.Equal(t, 3, q.RomanToInt("III"))
	})

	t.Run("test 2", func(t *testing.T) {
		assert.Equal(t, 58, q.RomanToInt("LVIII"))
	})

	t.Run("test 3", func(t *testing.T) {
		assert.Equal(t, 1994, q.RomanToInt("MCMXCIV"))
	})
}
