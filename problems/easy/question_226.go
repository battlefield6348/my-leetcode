package easy

type Question226 struct{}

func NewQuestion226() *Question226 {
	return &Question226{}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (q Question226) InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	new := &TreeNode{}
	new.Val = root.Val
	new.Left = root.Left
	new.Right = root.Right

	root.Left = q.InvertTree(new.Right)
	root.Right = q.InvertTree(new.Left)
	return root
}
