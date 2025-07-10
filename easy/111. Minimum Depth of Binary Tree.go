package easy

/*
type Pair struct {
	node *TreeNode
	dist int
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
} */

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []Pair{{root, 1}}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr.node.Left == nil && curr.node.Right == nil {
			return curr.dist
		}
		if curr.node.Left != nil {
			queue = append(queue, Pair{curr.node.Left, curr.dist + 1})
		}
		if curr.node.Right != nil {
			queue = append(queue, Pair{curr.node.Right, curr.dist + 1})
		}
	}
	return 0
}
