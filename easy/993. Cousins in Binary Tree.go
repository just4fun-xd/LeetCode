package easy

type Trinity struct {
	node   *TreeNode
	parent *TreeNode
	dist   int
}

func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil {
		return false
	}

	queue := []Trinity{{root, nil, 1}}
	parents := make(map[int]Trinity)
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		val := curr.node.Val
		if val == x || val == y {
			if pair, ok := parents[x+y-val]; ok {
				return pair.parent != curr.parent && pair.dist == curr.dist
			}
			parents[val] = curr
		}

		if curr.node.Left != nil {
			queue = append(queue, Trinity{curr.node.Left, curr.node, curr.dist + 1})
		}
		if curr.node.Right != nil {
			queue = append(queue, Trinity{curr.node.Right, curr.node, curr.dist + 1})
		}
	}
	return false
}
