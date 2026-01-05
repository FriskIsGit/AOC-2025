package days

import (
	"aoc-2025/util"
	"math"
	"strings"
)

// In Day 9's input, each line represents 0-indexed coordinates in the following format: col,row

type Location []int

func (l Location) Area(other Location) uint64 {
	col, row := l[0], l[1]
	otherCol, otherRow := other[0], other[1]
	width := uint64(util.AbsInt(col-otherCol) + 1)
	height := uint64(util.AbsInt(row-otherRow) + 1)
	return width * height
}

func Day9Part1(lines []string) uint64 {
	reds := make([]Location, 0, len(lines))
	for _, line := range lines {
		split := strings.Split(line, ",")
		if len(split) != 2 {
			continue
		}
		col, _ := util.ParseInt(split[0])
		row, _ := util.ParseInt(split[1])
		reds = append(reds, Location{col, row})
	}
	/*minCol, minRow := math.MaxInt, math.MaxInt
	maxCol, maxRow := -1, -1
	for _, red := range reds {
		minCol = min(red[0], minCol)
		minRow = min(red[1], minRow)
		maxCol = max(red[0], maxCol)
		maxRow = max(red[1], maxRow)
	}
	maxArea := uint64(0)
	topMin, topMax := getEdgeTiles(reds, HORIZONTAL, minRow)
	bottomMin, bottomMax := getEdgeTiles(reds, HORIZONTAL, maxRow)
	leftMin, leftMax := getEdgeTiles(reds, VERTICAL, minCol)
	rightMin, rightMax := getEdgeTiles(reds, VERTICAL, maxCol)

	edgePoints := []Location{topMin, topMax, bottomMin, bottomMax, leftMin, leftMax, rightMin, rightMax}

	for i := 0; i < len(edgePoints)-1; i++ {
		for j := i + 1; j < len(edgePoints); j++ {
			area := edgePoints[i].Area(edgePoints[j])
			maxArea = max(maxArea, area)
		}
	}*/
	return Day9Part1BruteForce(reds)
}

func Day9Part1BruteForce(reds []Location) uint64 {
	maxArea := uint64(0)
	for i := 0; i < len(reds)-1; i++ {
		for j := i + 1; j < len(reds); j++ {
			area := reds[i].Area(reds[j])
			maxArea = max(maxArea, area)
		}
	}
	return maxArea
}

type Orientation int

const (
	HORIZONTAL Orientation = iota
	VERTICAL
)

func getEdgeTiles(reds []Location, orientation Orientation, onIndex int) (Location, Location) {
	minIndex, maxIndex := math.MaxInt, -1
	var minLocation, maxLocation Location
	for _, red := range reds {
		col, row := red[0], red[1]
		if orientation == HORIZONTAL {
			if row != onIndex {
				continue
			}
			if col < minIndex {
				minIndex = col
				minLocation = red
			}
			if col > maxIndex {
				maxIndex = col
				maxLocation = red
			}
		} else {
			if col != onIndex {
				continue
			}
			if row < minIndex {
				minIndex = row
				minLocation = red
			}
			if row > maxIndex {
				maxIndex = row
				maxLocation = red
			}
		}

	}
	return minLocation, maxLocation
}
