package range_sum_of_BST_938

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sum int

func rangeSumBST(root *TreeNode, low int, high int) int {
	sum = 0
	return helper(root, low, high)
}

func helper(root *TreeNode, low, high int) int {
	if root == nil {
		return 0
	}

	if root.Val >= low && root.Val <= high {
		sum = sum + root.Val
		helper(root.Right, low, high)
		helper(root.Left, low, high)
	}
	if root.Val < low {
		helper(root.Right, low, high)
	}
	if root.Val > high {
		helper(root.Left, low, high)
	}
	return sum
}
