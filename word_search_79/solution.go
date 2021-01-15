package word_search_79

func wordsearch(board [][]string, word string) bool {
	index := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if dfs(board, i, j, word, index) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]string, i, j int, word string, index int) bool {
	if index == len(word) {
		return true
	}
	if index > len(word) || i < 0 || j < 0 || i == len(board) || j == len(board[i]) || board[i][j] != string(word[index]) {
		return false
	}
	c := board[i][j]
	board[i][j] = "#"
	help := [4][2]int{{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1}}

	for k := 0; k < 4; k++ {
		if dfs(board, i+help[k][0], j+help[k][1], word, index+1) {
			return true
		}
	}
	board[i][j] = c
	return false
}
