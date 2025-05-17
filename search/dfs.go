package search

// example: give a word, and search this word in board to check whether exist or not
// prePos: the previous one position, could not be used again normally.
// marks: recording used or not
// 1. check current situation: current match??? visited??? all match???
// 2. work through next situation: do not go back, watch out the index of position
func dfs(board [][]byte, visitedMark [][]bool, currPos, target []int, currSteps int, minSteps *int) {
	currX := currPos[0]
	currY := currPos[1]
	if visitedMark[currX][currY] {
		return
	}

	// set mark, this position can not be used
	visitedMark[currX][currY] = true
	// recover after processed
	defer func(visitedMark [][]bool, x, y int) {
		visitedMark[x][y] = false
	}(visitedMark, currX, currY)

	// reach the target point
	if currX == target[0] && currY == target[1] {
		if currSteps < *minSteps {
			*minSteps = currSteps
		}
		return
	}

	posDiff := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, diff := range posDiff {
		newX := currX + diff[0]
		newY := currY + diff[1]
		// out of index
		if newX < 0 || newX >= len(board) || newY < 0 || newY >= len(board[0]) {
			continue
		}

		// never go back
		if visitedMark[newX][newY] {
			continue
		}

		// skip block points
		if board[newX][newY] != 1 {
			continue
		}

		// go next, no need to set mark, was setted after step into this function
		dfs(board, visitedMark, []int{newX, newY}, target, currSteps+1, minSteps)
		// found, return is ok
	}
}

func Dfs(board [][]byte, startPoint, endPoint []int) int {
	result := -1
	if len(startPoint) != 2 || len(endPoint) != 2 {
		return result
	}
	lenX, lenY := len(board), len(board[0])
	if startPoint[0] < 0 || startPoint[0] >= lenX || startPoint[1] < 0 || startPoint[1] >= lenY {
		return result
	}
	if endPoint[0] < 0 || endPoint[0] >= lenX || endPoint[1] < 0 || endPoint[1] >= lenY {
		return result
	}

	visitedMark := make([][]bool, len(board))
	for i := 0; i < lenX; i++ {
		if len(board[i]) != lenY {
			return result
		}
		visitedMark[i] = make([]bool, lenY)
		for j := 0; j < lenY; j++ {
			visitedMark[i][j] = false
		}
	}

	steps := 9999
	dfs(board, visitedMark, startPoint, endPoint, 0, &steps)
	if steps == 9999 {
		return -1
	}
	return steps
}
