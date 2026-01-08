package days

import (
	"aoc-2025/util"
	"math"
	"strings"
)

// In Day 9's input, each line represents 0-indexed coordinates in the following format: col,row

type Point []int

func (l Point) Area(other Point) uint64 {
	col, row := l[0], l[1]
	otherCol, otherRow := other[0], other[1]
	width := uint64(util.AbsInt(col-otherCol) + 1)
	height := uint64(util.AbsInt(row-otherRow) + 1)
	return width * height
}

func Day9Part1(lines []string) uint64 {
	reds := make([]Point, 0, len(lines))
	for _, line := range lines {
		split := strings.Split(line, ",")
		if len(split) != 2 {
			continue
		}
		col, _ := util.ParseInt(split[0])
		row, _ := util.ParseInt(split[1])
		reds = append(reds, Point{col, row})
	}
	return Day9Part1BruteForce(reds)
}

func Day9Part2(lines []string) uint64 {
	corners := make([]Point, 0, len(lines))
	// Check if some points are not actually corners but reside on a line
	for _, line := range lines {
		split := strings.Split(line, ",")
		if len(split) != 2 {
			continue
		}
		col, _ := util.ParseInt(split[0])
		row, _ := util.ParseInt(split[1])
		corners = append(corners, Point{col, row})
	}
	maxArea := uint64(0)
	for i := 0; i < len(corners)-1; i++ {
		for j := i + 1; j < len(corners); j++ {
			area := corners[i].Area(corners[j])
			if area > maxArea && isFullyContained(corners[i], corners[j], corners) {
				maxArea = area
			}
		}
	}
	return Day9Part1BruteForce(corners)
}

func isFullyContained(corner, corner2 Point, corners []Point) bool {
	corner3 := Point{corner[0], corner2[1]}
	corner4 := Point{corner2[0], corner[1]}
	_ = corner3
	_ = corner4
	return false
}

func Day9Part1BruteForce(reds []Point) uint64 {
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

func getEdgeTiles(reds []Point, orientation Orientation, onIndex int) (Point, Point) {
	minIndex, maxIndex := math.MaxInt, -1
	var minLocation, maxLocation Point
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
