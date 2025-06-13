package easy_test

import (
	"my-leetcode/problems/easy"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuestion1(t *testing.T) {
	q := easy.NewQuestion1()
	t.Run("test 1", func(t *testing.T) {
		assert.Equal(t, []int{0, 1}, q.TwoSum([]int{2, 7, 11, 15}, 9))
	})

	t.Run("test 2", func(t *testing.T) {
		assert.Equal(t, []int{1, 2}, q.TwoSum([]int{3, 2, 4}, 6))
	})

	t.Run("test 3", func(t *testing.T) {
		assert.Equal(t, []int{0, 2}, q.TwoSum([]int{-3, 4, 3, 90}, 0))
	})
}
