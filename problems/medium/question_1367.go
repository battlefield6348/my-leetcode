package medium

type Question1367 struct{}

func NewQuestion1367() Question1367 {
	return Question1367{}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/linked-list-in-binary-tree/description/?envType=daily-question&envId=2024-09-07
func (q *Question1367) Question1367(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}

	if root == nil {
		return false
	}

	if head.Val == root.Val {
		if q.Check(head, root) {
			return true
		}

	}
	return q.Question1367(head, root.Left) || q.Question1367(head, root.Right)
}

func (q *Question1367) Check(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}

	if root == nil {
		return false
	}

	if root.Val == head.Val {
		return q.Check(head.Next, root.Left) || q.Check(head.Next, root.Right)
	}
	return false
}

func (q *Question1367) GenerateListNode(list []int) *ListNode {
	if len(list) == 0 {
		return nil
	}

	dummy := &ListNode{}
	current := dummy

	for _, val := range list {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}

	return dummy.Next // 返回虛擬頭節點的下一個節點，即實際的頭節點
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (q *Question1367) GenerateTreeNode(treeNodeList map[int]*int) *TreeNode {
	if val, ok := treeNodeList[1]; ok && val != nil {
		root := &TreeNode{Val: *val}
		q.buildTree(root, treeNodeList, 1)
		return root
	}
	return nil
}

func (q *Question1367) buildTree(node *TreeNode, m map[int]*int, index int) {
	leftIndex := 2 * index
	rightIndex := 2*index + 1

	if left, ok := m[leftIndex]; ok && left != nil {
		node.Left = &TreeNode{Val: *left}
		q.buildTree(node.Left, m, leftIndex)
	}

	if right, ok := m[rightIndex]; ok && right != nil {
		node.Right = &TreeNode{Val: *right}
		q.buildTree(node.Right, m, rightIndex)
	}
}
