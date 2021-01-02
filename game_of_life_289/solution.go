package game_of_life_289

func gameOfLife(board [][]int) [][]int {
	if len(board) == 0 {
		return board
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			liveN := getLive(i, j, board)
			if board[i][j] == 1 && (liveN < 2 || liveN > 3) {
				board[i][j] = 2
			} else if board[i][j] == 0 && liveN == 3 {
				board[i][j] = -1
			}
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 2 {
				board[i][j] = 0
			} else if board[i][j] == -1 {
				board[i][j] = 1
			}
		}
	}
	return board
}

func getLive(i, j int, b [][]int) int {
	chi := []int{1, -1, 1, -1, -1, 0, 1, 0}
	chj := []int{1, -1, -1, 1, 0, 1, 0, -1}
	liveN := 0
	for ch := 0; ch < 8; ch++ {
		if valid(i+chi[ch], j+chj[ch], len(b), len(b[0])) && (b[i+chi[ch]][j+chj[ch]] == 1 || b[i+chi[ch]][j+chj[ch]] == 2) {
			liveN++
		}
	}
	return liveN
}
func valid(i, j, n, m int) bool {
	if i < 0 || i >= n || j < 0 || j >= m {
		return false
	}
	return true
}
