package easy_test

import (
	"my-leetcode/problems/easy"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuestion226(t *testing.T) {
	q := easy.NewQuestion226()

	t.Run("test 1", func(t *testing.T) {
		ans := q.InvertTree(&easy.TreeNode{
			Val: 4,
			Left: &easy.TreeNode{
				Val: 2,
				Left: &easy.TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
				Right: &easy.TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &easy.TreeNode{
				Val: 7,
				Left: &easy.TreeNode{
					Val:   6,
					Left:  nil,
					Right: nil,
				},
				Right: &easy.TreeNode{
					Val:   9,
					Left:  nil,
					Right: nil,
				},
			},
		})
		assert.Equal(t, 4, ans.Val)
		assert.Equal(t, 7, ans.Left.Val)
		assert.Equal(t, 2, ans.Right.Val)
		assert.Equal(t, 9, ans.Left.Left.Val)
		assert.Equal(t, 6, ans.Left.Right.Val)
		assert.Equal(t, 3, ans.Right.Left.Val)
		assert.Equal(t, 1, ans.Right.Right.Val)
	})

	t.Run("test 2", func(t *testing.T) {
		ans := q.InvertTree(&easy.TreeNode{
			Val: 2,
			Left: &easy.TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &easy.TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		})

		assert.Equal(t, 2, ans.Val)
		assert.Equal(t, 3, ans.Left.Val)
		assert.Equal(t, 1, ans.Right.Val)
	})
}
