package days

import (
	"aoc-2025/util"
	"fmt"
)

// --- Day 7: Laboratories ---

type Beam []int

func Day7Part1(lines []string) int {
	lines = removeEmptySpaceLines(lines)
	board := util.ToBoard(lines)
	startCol := findStartPointColumn(lines)
	splits := 0
	var beams []Beam
	beams = append(beams, Beam{0, startCol})
	var allBeams []Beam

	rowCount := len(lines)
	width := len(lines[0])
	for len(beams) > 0 {
		lastIndex := len(beams) - 1
		rootBeam := beams[lastIndex]
		row := rootBeam[0]
		col := rootBeam[1]
		beams = util.DeleteAt(beams, lastIndex)

		for r := row + 1; r < rowCount; r++ {
			cell := lines[r][col]
			if cell != '^' {
				continue
			}
			fmt.Println()
			fmt.Println()
			util.PrintBoard(board)
			splits++
			// Splitters are never right next to each other so it's safe to spawn beams on splitter's sides
			leftCol := col - 1
			rightCol := col + 1
			if leftCol >= 0 {
				board[r][leftCol] = '|'
				leftBeam := Beam{r, leftCol}
				if !containsBeam(allBeams, leftBeam) {
					beams = append(beams, leftBeam)
					allBeams = append(allBeams, leftBeam)
				}
			}
			if rightCol < width {
				board[r][rightCol] = '|'
				rightBeam := Beam{r, rightCol}
				if !containsBeam(allBeams, rightBeam) {
					beams = append(beams, rightBeam)
					allBeams = append(allBeams, rightBeam)
				}
			}
			break
		}
	}
	util.PrintBoard(board)
	return splits
}

func containsBeam(beams []Beam, beam Beam) bool {
	for _, b := range beams {
		if b[0] == beam[0] && b[1] == beam[1] {
			return true
		}
	}
	return false
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

func findStartPointColumn(board []string) int {
	row0 := board[0]
	for c, val := range row0 {
		if val == 'S' {
			return c
		}
	}
	return -1
}
