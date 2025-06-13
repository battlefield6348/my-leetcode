package easy_test

import (
	"my-leetcode/problems/easy"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuestion21(t *testing.T) {
	q := easy.NewQuestion21()
	t.Run("test 1", func(t *testing.T) {
		listNode := q.MergeTwoLists(
			&easy.ListNode{Val: 1, Next: &easy.ListNode{Val: 2, Next: &easy.ListNode{Val: 4}}},
			&easy.ListNode{Val: 1, Next: &easy.ListNode{Val: 3, Next: &easy.ListNode{Val: 4}}},
		)
		assert.Equal(t, 1, listNode.Val)
		assert.Equal(t, 1, listNode.Next.Val)
		assert.Equal(t, 2, listNode.Next.Next.Val)
		assert.Equal(t, 3, listNode.Next.Next.Next.Val)
		assert.Equal(t, 4, listNode.Next.Next.Next.Next.Val)
		assert.Equal(t, 4, listNode.Next.Next.Next.Next.Next.Val)
	})
}
