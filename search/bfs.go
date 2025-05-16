package search

const (
	MAX_STEP_COUNT = 99999
)

// 迷宫中从某一点到另一点的最短路径，迷宫中存在一些不可走的点
// normal question: find the min step to target position, no need to record path info
func Bfs(board [][]int, endPos []int) int {
	// check param: board
	lenX, lenY := len(board), len(board[0])
	for i := 0; i < lenX; i++ {
		// length of board[i] must be equal
		if len(board[i]) != lenY {
			return -1
		}
		// only -1 and 1 will be accepted
		for j := 0; j < lenY; j++ {
			if board[i][j] != 1 && board[i][j] != -1 {
				return -1
			}
		}
	}

	// check param: endPos
	if len(endPos) != 2 {
		return -1
	}
	// out of board
	if endPos[0] < 0 || endPos[0] >= lenX {
		return -1
	}
	if endPos[1] < 0 || endPos[1] >= lenY {
		return -1
	}

	// unreachable point
	if board[endPos[0]][endPos[1]] == -1 {
		return -1
	}

	// build a array, recording the min steps to specify point
	startPos := []int{0, 0}
	visitedCount := make([][]int, lenX)
	for i := 0; i < lenX; i++ {
		visitedCount[i] = make([]int, lenY)
		for j := range board[i] {
			visitedCount[i][j] = MAX_STEP_COUNT
		}
	}
	visitedCount[startPos[0]][startPos[1]] = 0

	// constant, to the next point
	posDiff := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	// storage the points waitting to be visited
	nextPos := [][]int{startPos}
	left, right := 0, 1
	for left < right {
		currX, currY := nextPos[left][0], nextPos[left][1]
		for _, diff := range posDiff {
			newX, newY := currX+diff[0], currY+diff[1]
			// new positon out of index
			if 0 > newX || newX >= lenX || 0 > newY || newY >= lenY {
				continue
			}

			// ignore if new position can not be used, such as block points
			if board[newX][newY] == -1 {
				continue
			}

			// new position was visited, and better than current
			if visitedCount[newX][newY] <= visitedCount[currX][currY]+1 {
				continue
			}

			visitedCount[newX][newY] = visitedCount[currX][currY] + 1
			// new position can be used, and need to be used
			nextPos = append(nextPos, []int{newX, newY})
			right++
		}
		left++
	}

	// unreachable point
	if visitedCount[endPos[0]][endPos[1]] == MAX_STEP_COUNT {
		return -1
	}

	// visitedCount[x][y] is the min step to point (x,y)
	return visitedCount[endPos[0]][endPos[1]]
}
