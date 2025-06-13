package easy

type Question21 struct{}

func NewQuestion21() *Question21 {
	return &Question21{}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (q *Question21) MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val <= list2.Val {
		return &ListNode{
			Val:  list1.Val,
			Next: q.MergeTwoLists(list1.Next, list2),
		}
	}
	return &ListNode{
		Val:  list2.Val,
		Next: q.MergeTwoLists(list1, list2.Next),
	}
}
