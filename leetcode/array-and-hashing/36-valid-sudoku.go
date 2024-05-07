package main

import "strconv"

func isValidSudokuWithoutOptimize(board [][]byte) bool {
	startRow := 0
	startCol := 0

	boarder := 9
	for startRow < boarder {

		numMap := make(map[byte]int)
		for i := startRow; i < startRow+3; i++ {
			for j := startCol; j < startCol+3; j++ {
				if board[i][j] != '.' {
					numMap[board[i][j]]++
					if numMap[board[i][j]] > 1 {
						return false
					}
				}
			}
		}

		startCol += 3
		if startCol >= boarder {
			startRow += 3
			startCol = 0
		}
	}

	//检查行
	startCol = 0
	for startCol < boarder {
		numMap := make(map[byte]int)
		for i := 0; i < boarder; i++ {
			if board[startCol][i] != '.' {
				numMap[board[startCol][i]]++
				if numMap[board[startCol][i]] > 1 {
					return false
				}
			}
		}
		startCol++
	}

	//检查列
	startRow = 0
	for startRow < boarder {
		numMap := make(map[byte]int)
		for i := 0; i < boarder; i++ {
			if board[i][startRow] != '.' {
				numMap[board[i][startRow]]++
				if numMap[board[i][startRow]] > 1 {
					return false
				}
			}
		}
		startRow++
	}

	return true
}

/*
	思路：
		我们需要检查的内容，9行，,9 列，9个3X3正方块, 只要将元素映射到对应的地方，完全可以一次遍历就把全部内容检查完成
		关于小九宫格映射，我们可以按照横向
		0，1，2，
		3，4，5，
		6，7，8
		来进行9个小的九宫格映射，第二行映射是第一行的值+3，也就是按照 列/3 * 3 相加进去即可
*/

func isValidSudoku(board [][]byte) bool {
	var columns, rows, grids = [9][9]bool{}, [9][9]bool{}, [9][9]bool{}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {

			val, err := strconv.Atoi(string(board[i][j]))
			if err != nil {
				continue
			}
			val--
			gridIndex := (i/3)*3 + j/3
			if columns[j][val] || rows[i][val] || grids[gridIndex][val] {
				return false
			}
			columns[j][val] = true
			rows[i][val] = true
			grids[gridIndex][val] = true
		}
	}
	return true
}
