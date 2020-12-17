package valid_binary_search_tree_98

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func help(root *TreeNode, min, max interface{}) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Val <= min.(int) || max != nil && root.Val >= max.(int) {
		return false
	}

	return help(root.Right, root.Val, max) && help(root.Left, min, root.Val)
}
func isValidBST(root *TreeNode) bool {
	return help(root, nil, nil)
}
