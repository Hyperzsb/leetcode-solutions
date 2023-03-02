/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

func construct(grid [][]int) *Node {
	root := Node{Val: true, IsLeaf: true,
		TopLeft: nil, TopRight: nil, BottomLeft: nil, BottomRight: nil}

	firstVal, isLeaf := grid[0][0], true
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != firstVal {
				isLeaf = false
				break
			}
		}
	}

	if isLeaf {
		if grid[0][0] == 0 {
			root.Val = false
		}

		return &root
	}

	root.IsLeaf = false

	topLeft := make([][]int, len(grid)/2)
	topRight := make([][]int, len(grid)/2)
	bottomLeft := make([][]int, len(grid)/2)
	bottomRight := make([][]int, len(grid)/2)
	for i := 0; i < len(grid)/2; i++ {
		topLeft[i] = make([]int, len(grid[0])/2)
		topRight[i] = make([]int, len(grid[0])/2)
		bottomLeft[i] = make([]int, len(grid[0])/2)
		bottomRight[i] = make([]int, len(grid[0])/2)

		for j := 0; j < len(grid[0])/2; j++ {
			topLeft[i][j] = grid[i][j]
			topRight[i][j] = grid[i][len(grid[0])/2+j]
			bottomLeft[i][j] = grid[len(grid)/2+i][j]
			bottomRight[i][j] = grid[len(grid)/2+i][len(grid[0])/2+j]
		}
	}

	root.TopLeft = construct(topLeft)
	root.TopRight = construct(topRight)
	root.BottomLeft = construct(bottomLeft)
	root.BottomRight = construct(bottomRight)

	return &root
}

