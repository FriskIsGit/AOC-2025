package days

import (
	"aoc-2025/util"
	"strings"
)

// --- Day 2: Gift Shop ---

func Day2Part1(input string) int {
	strRanges := strings.Split(input, ",")
	invalidSum := 0
	var ranges []Range
	for _, strRange := range strRanges {
		r := ParseRange(strRange)
		ranges = append(ranges, r)
	}
	return invalidSum
}

type Range struct {
	start int64
	end   int64
}

func newRange(start, end int64) Range {
	return Range{start, end}
}

func ParseRange(strRange string) Range {
	left, right, _ := strings.Cut(strRange, "-")
	start, _ := util.ParseLong(left)
	end, _ := util.ParseLong(right)
	return newRange(start, end)
}
