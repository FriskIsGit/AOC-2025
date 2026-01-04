package days

import (
	"aoc-2025/util"
)

// --- Day 7: Laboratories ---

type Beam []int

func Day7Part1(lines []string) int {
	lines = removeEmptySpaceLines(lines)
	board := util.ToBoard(lines)
	startCol := findStartPointColumn(lines[0])
	splits := 0
	var beams []Beam
	beams = append(beams, Beam{1, startCol})

	rowCount := len(lines)
	width := len(lines[0])
	for len(beams) > 0 {
		lastIndex := len(beams) - 1
		rootBeam := beams[lastIndex]
		row, col := rootBeam[0], rootBeam[1]
		beams = util.DeleteAt(beams, lastIndex)

		for r := row; r < rowCount; r++ {
			cell := board[r][col]
			if cell == '|' {
				break
			}
			if cell == '.' {
				board[r][col] = '|'
				continue
			}
			// Must be a splitter
			splits++
			// Splitters are never right next to each other so it's safe to spawn beams on splitter's sides
			leftCol := col - 1
			rightCol := col + 1
			if leftCol >= 0 {
				leftBeam := Beam{r, leftCol}
				beams = append(beams, leftBeam)
			}
			if rightCol < width {
				rightBeam := Beam{r, rightCol}
				beams = append(beams, rightBeam)
			}
			break
		}
	}
	util.PrintBoard(board)
	return splits
}

func Day7Part2(lines []string) int {
	lines = removeEmptySpaceLines(lines)
	board := util.ToBoard(lines)
	startCol := findStartPointColumn(lines[0])
	memoBoard := createMemoBoard(board)
	width := len(board[0])
	height := len(board)
	for r := height - 1; r >= 0; r-- {
		for c := 0; c < width; c++ {
			if board[r][c] != '^' {
				continue
			}
			leftCol, rightCol := c-1, c+1
			leftWays, rightWays := 0, 0
			if leftCol >= 0 {
				leftWays = countTimelines(r, leftCol, board, memoBoard)
			}
			if rightCol < width {
				rightWays = countTimelines(r, rightCol, board, memoBoard)
			}
			ways := leftWays + rightWays
			memoBoard[r][c] = ways
		}
	}
	return countTimelines(0, startCol, board, memoBoard)
}

func countTimelines(row, col int, board [][]byte, memoBoard [][]int) int {
	height := len(board)
	// Position row below the original splitter
	for down := row + 1; down < height; down++ {
		if board[down][col] == '^' {
			return memoBoard[down][col]
		}
	}
	return 1
}

func createMemoBoard(board [][]byte) [][]int {
	memoBoard := make([][]int, len(board))
	for i, row := range board {
		memoBoard[i] = make([]int, len(row))
	}
	return memoBoard
}

func removeEmptySpaceLines(lines []string) []string {
	capacity := 2 + len(lines)/2
	filtered := make([]string, 0, capacity)

	take := true
	for _, line := range lines {
		if take {
			filtered = append(filtered, line)
		}
		take = !take
	}
	return filtered
}

func findStartPointColumn(row string) int {
	for c, val := range row {
		if val == 'S' {
			return c
		}
	}
	return -1
}
