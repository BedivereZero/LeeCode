package algorithms

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	// n: 0, 1
	switch n {
	case 0:
		return nil
	case 1:
		return []*TreeNode{
			{
				Val: 1,
			},
		}
	}

	// Create 3D matrix: cache
	cache := make([][][]*TreeNode, n+1)
	for x := 0; x < n+1; x++ {
		cache[x] = make([][]*TreeNode, n+2-x)
	}

	// Initial 0, 1 column
	for x := 0; x < n+1; x++ {
		cache[x][0] = []*TreeNode{nil}
	}
	for x := 0; x < n; x++ {
		cache[x][1] = []*TreeNode{
			{
				Val: x + 1,
			},
		}
	}

	// Calculate reset of cache
	for y := 2; y < n+2; y++ {
		for x := 0; x < n+1-y; x++ {
			for i := 0; i < y; i++ {
				// cache[x][i] * cache[x + 1 + i][y - 1 - i]
				for _, l := range cache[x][i] {
					for _, r := range cache[x+1+i][y-1-i] {
						cache[x][y] = append(cache[x][y], &TreeNode{
							Val:   x + i + 1,
							Left:  copyTree(l),
							Right: copyTree(r),
						})
					}
				}
			}
		}
	}
	return cache[0][n]
}

func copyTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	return &TreeNode{
		Val:   root.Val,
		Left:  copyTree(root.Left),
		Right: copyTree(root.Right),
	}
}
