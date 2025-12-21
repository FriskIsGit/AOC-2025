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

func Day1Part2(lines []string) int {
	dialPos := 50
	zeroHits := 0
	for _, line := range lines {
		dir := line[0]
		turns, _ := util.ParseInt(line[1:])
		var nextPos, hits int
		if dir == 'L' {
			hits = GetZeroHits(dialPos, -turns)
			nextPos = NextDialPos(dialPos, -turns)
		} else {
			hits = GetZeroHits(dialPos, turns)
			nextPos = NextDialPos(dialPos, turns)
		}
		zeroHits += hits
		dialPos = nextPos
	}

	return zeroHits
}

func NextDialPos(currentPos int, turns int) int {
	nextPos := currentPos + turns
	if nextPos < 0 {
		rem := nextPos % 100
		if rem == 0 {
			nextPos = 0
		} else {
			nextPos = 100 + rem
		}
	} else if nextPos > 99 {
		rem := nextPos % 100
		nextPos = rem
	}
	return nextPos
}

func GetZeroHits(currentPos int, turns int) int {
	hits := turns / 100
	if turns < 0 {
		hits = -hits
		turns = turns + (hits * 100)
	} else {
		turns = turns - (hits * 100)
	}
	// If no leftover turns or currentPos == 0 then there's no possible move that'd cross zero
	if turns == 0 || currentPos == 0 {
		return hits
	}
	finalPos := currentPos + turns
	if finalPos <= 0 || finalPos >= 100 {
		hits++
	}
	return hits
}
