package main

// 回溯算法 - 八皇后问题

func BackTrace(n int) {
	board := chessBoard(n)
	r := 0
	for i := 0; i < n; i++ {
		board[r][i] = 1
	}

}

func chessBoard(n int) [][]int {
	board := make([][]int, n)
	for i := 0; i < len(board); i++ {
		board[i] = make([]int, n)
	}
	return board
}

func main() {
	BackTrace(8)
}
