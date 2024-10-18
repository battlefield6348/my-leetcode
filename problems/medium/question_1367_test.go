package medium_test

import (
	"my-leetcode/problems/medium"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuestion1367(t *testing.T) {
	q := medium.NewQuestion1367()

	t.Run("test 1", func(t *testing.T) {
		head := q.GenerateListNode([]int{4, 2, 8})
		tree := q.GenerateTreeNode(map[int]*int{
			1:  intPrt(1),
			2:  intPrt(4),
			3:  intPrt(4),
			4:  nil,
			5:  intPrt(2),
			6:  intPrt(2),
			7:  nil,
			8:  intPrt(1),
			9:  nil,
			10: intPrt(6),
			11: intPrt(8),
			12: nil,
			13: nil,
			14: nil,
			15: nil,
			16: nil,
			17: intPrt(1),
			18: intPrt(3),
		})

		ans := q.Question1367(head, tree)
		assert.True(t, ans)
	})

	// t.Run("test 2", func(t *testing.T) {
	// 	ans := q.Question1367("leetcode", 2)
	// })

	// t.Run("test 3", func(t *testing.T) {
	// 	ans := q.Question1367("zbax", 2)
	// })

	// t.Run("test 4", func(t *testing.T) {
	// 	ans := q.Question1367("vbyytoijnbgtyrjlsc", 2)
	// 	assert.Equal(t, 15, ans)
	// })
}

func TestGenerateListNode(t *testing.T) {
	q := medium.NewQuestion1367()

	t.Run("test 1", func(t *testing.T) {
		nodeList := q.GenerateListNode([]int{4})
		assert.Equal(t, &medium.ListNode{
			Val:  4,
			Next: nil,
		}, nodeList)
	})

	t.Run("test 2", func(t *testing.T) {
		nodeList := q.GenerateListNode([]int{4, 2})
		assert.Equal(t, &medium.ListNode{
			Val: 4,
			Next: &medium.ListNode{
				Val:  2,
				Next: nil,
			},
		}, nodeList)
	})

	t.Run("test 3", func(t *testing.T) {
		nodeList := q.GenerateListNode([]int{4, 2, 8, 7})
		assert.Equal(t, &medium.ListNode{
			Val: 4,
			Next: &medium.ListNode{
				Val: 2,
				Next: &medium.ListNode{
					Val: 8,
					Next: &medium.ListNode{
						Val:  7,
						Next: nil,
					}},
			},
		}, nodeList)
	})
}

func intPrt(i int) *int {
	return &i
}

func TestGenerateTreeNode(t *testing.T) {
	q := medium.NewQuestion1367()

	t.Run("test 1", func(t *testing.T) {
		treeNode := q.GenerateTreeNode(map[int]*int{
			1: intPrt(1),
			2: intPrt(4),
			3: intPrt(4),
			4: nil,
			5: intPrt(2),
			6: intPrt(2),
			7: nil,
		})
		assert.Equal(t, &medium.TreeNode{
			Val: 1,
			Left: &medium.TreeNode{
				Val: 4,
				Right: &medium.TreeNode{
					Val: 2,
				},
			},
			Right: &medium.TreeNode{
				Val: 4,
				Left: &medium.TreeNode{
					Val: 2,
				},
			},
		}, treeNode)
	})

	t.Run("test 2", func(t *testing.T) {
		treeNode := q.GenerateTreeNode(map[int]*int{
			1: intPrt(1),
		})

		assert.Equal(t, &medium.TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		}, treeNode)
	})
}
