package days

import "aoc-2025/util"

// --- Day 4: Printing Department ---

func Day4Part1(lines []string) int {
	board := util.ToBoard(lines)
	rows := len(board)
	width := len(board[0])
	paperRolls := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < width; c++ {
			if board[r][c] == '.' {
				continue
			}
			count := CountAdjacent('@', board, r, c)
			if count < 4 {
				paperRolls++
			}
		}
	}
	return paperRolls
}

func Day4Part2(lines []string) int {
	board := util.ToBoard(lines)
	rows := len(board)
	width := len(board[0])
	paperRolls := 0
	for {
		rollsBefore := paperRolls
		for r := 0; r < rows; r++ {
			for c := 0; c < width; c++ {
				if board[r][c] == '.' {
					continue
				}
				count := CountAdjacent('@', board, r, c)
				if count < 4 {
					board[r][c] = '.'
					paperRolls++
				}
			}
		}
		// No changes occurred - quit
		if paperRolls == rollsBefore {
			break
		}
	}

	return paperRolls
}

var ADJACENT_DISPLACEMENTS = [8][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func CountAdjacent(target byte, board [][]byte, row, col int) int {
	rows := len(board)
	width := len(board[0])
	count := 0
	for _, d := range ADJACENT_DISPLACEMENTS {
		dRow, dCol := row+d[0], col+d[1]
		if dRow >= 0 && dRow < rows && dCol >= 0 && dCol < width {
			if board[dRow][dCol] == target {
				count++
			}
		}
	}
	return count
}
