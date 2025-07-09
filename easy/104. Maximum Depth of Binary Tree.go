package easy

type Pair struct {
	node *TreeNode
	dist int
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []Pair{{root, 1}}
	maxDist := 1

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.dist > maxDist {
			maxDist = current.dist
		}

		if current.node.Left != nil {
			queue = append(queue, Pair{current.node.Left, current.dist + 1})
		}
		if current.node.Right != nil {
			queue = append(queue, Pair{current.node.Right, current.dist + 1})
		}

	}
	return maxDist
}
