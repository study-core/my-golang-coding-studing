package main

import "fmt"

func main() {

	findGraph(4)
}

// 回溯算法 (八 王后问题)

/**
这个问题很经典了，简单解释一下：给你一个 N×N 的棋盘，让你放置 N 个皇后，使得它们不能互相攻击。

PS：皇后可以攻击同一行、同一列、左上左下右上右下四个方向的任意单位
*/

// 解题思路:
// 		决策树的每一层表示棋盘上的每一行；每个节点可以做出的选择是，在该行的任意一列放置一个皇后.

/* 输入棋盘边长 n，返回所有合法的放置 */
func findGraph(n int) { // n 为 n位王后在 n * n 的棋盘上的布局

	// 初始化 空棋盘
	board := make([][]string, n)

	// 遍历 棋盘的 行
	for board_row := 0; board_row < len(board); board_row++ {

		rows := make([]string, n)
		// 每一行，遍历每个空格，填充 "." 表示 棋盘上的空格子
		for col := 0; col < len(rows); col++ {
			rows[col] = "."
		}

		// 将 初始化完毕的 行 填充棋盘
		board[board_row] = rows
	}

	fmt.Println("初始化完棋盘后")
	printBoard(board)

	// todo 开始 递归 进入回溯
	backtrack(board, 0)

	printBoard(board)
}

func printBoard (board [][]string) {
	//fmt.Println(board)
	for row := 0; row < len(board); row++ {
		fmt.Println(board[row])
	}
}

// todo 回溯
func backtrack(board [][]string, row int) {

	fmt.Println("进来操作第 ", row, " 行.......")

	// 路径：board 中小于 row 的那些行都已经成功放置了皇后
	// 选择列表：第 row 行的所有列都是放置皇后的选择
	// 结束条件：row 超过 board 的最后一行

	// todo 触发结束条件
	if row == len(board) - 1 {
		fmt.Println("结束遍历, 第 ", row, " 行.......")
		return // todo 返回最后棋盘的布局
	}

	// todo 在当前行, 遍历各个 小格子
	colNum := len(board[row])

	for col := 0; col < colNum; col++ {
		fmt.Println("开始遍历第 ", row, " 行的, 列数, 第 ", col, " 列")
		// todo 排除不合法选择
		if !isValid(board, row, col) {
			continue
		}
		// todo 做选择,   当前空格放置王后 `Q`
		board[row][col] = "Q"
		fmt.Println("【开始放】, 第 ", row, " 行的, 列数, 第 ", col, " 列 的 Q")


		// todo 递归 进入下一行决策
		backtrack(board, row+1)

		// todo 撤销选择, 如果 递归能够出来, 说明已经走到头了, 需要 退出来 回溯上一步
		//
		// 还原该步 棋盘
		board[row][col] = "."
		fmt.Println("【清空】, 第 ", row, " 行的, 列数, 第 ", col, " 列 的 Q")
	}
}

/* 是否可以在 board[row][col] 放置皇后？ */
func isValid(board [][]string, row, col int) bool {

	rowNum := len(board)

	// todo 逐行检查
	// 检查 ·列· 是否有皇后互相冲突
	for i := 0; i < rowNum; i++ {
		if board[i][col] == "Q" {

			fmt.Println( "row:", row, "当前row:", i, "col:", col, "发现【当前列】已经有 Q")
			printBoard(board)
			return false
		}
	}

	// 检查右上方是否有皇后互相冲突

	i := row - 1
	j := col + 1
	for i >= 0 && j < rowNum {
		if board[i][j] == "Q" {
			fmt.Println("row:", row, "当前row:", i, "col:", col, "发现【右上方】已经有 Q")
			printBoard(board)
			return false
		}
		i--
		j++
	}

	i = row - 1
	j = col - 1
	// 检查左上方是否有皇后互相冲突
	for i >= 0 && j >= 0 {
		if board[i][j] == "Q" {

			fmt.Println("row:", row, "当前row:", i, "col:", col, "发现 【左上方】已经有 Q")
			printBoard(board)
			return false
		}
		i--
		j--
	}
	return true
}
