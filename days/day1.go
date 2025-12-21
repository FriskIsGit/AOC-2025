package days

import "aoc-2025/util"

// --- Day 1: Secret Entrance ---

func Day1Part1(lines []string) int {
	dialPos := 50
	zeroHits := 0
	for _, line := range lines {
		dir := line[0]
		turns, _ := util.ParseInt(line[1:])
		var nextPos int
		if dir == 'L' {
			nextPos = NextDialPos(dialPos, -turns)
		} else {
			nextPos = NextDialPos(dialPos, turns)
		}
		if nextPos == 0 {
			zeroHits++
		}
		dialPos = nextPos
	}

	return zeroHits
}

func NextDialPos(currentPos int, turns int) int {
	nextPos := currentPos + turns
	if nextPos < 0 {
		rem := nextPos % 100
		nextPos = 100 + rem
	} else if nextPos > 99 {
		rem := nextPos % 100
		nextPos = rem
	}
	return nextPos
}
