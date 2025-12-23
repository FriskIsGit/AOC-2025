package days

import (
	"aoc-2025/util"
)

// --- Day 5: Cafeteria ---

func Day5Part1(lines []string) int {
	l := 0
	var ranges []util.Range
	for {
		if lines[l] == "" {
			break
		}
		rang := util.ParseRange(lines[l])
		ranges = append(ranges, rang)
		l++
	}

	var ids []int64
	for i := l + 1; i < len(lines); i++ {
		id, _ := util.ParseLong(lines[i])
		ids = append(ids, id)
	}

	fresh := 0
	for _, id := range ids {
		for _, rang := range ranges {
			if rang.Includes(id) {
				fresh++
				break
			}
		}
	}
	return fresh
}

func Day5Part2(lines []string) int64 {
	var ranges []util.Range
	for _, line := range lines {
		if line == "" {
			break
		}
		rang := util.ParseRange(line)
		ranges = append(ranges, rang)
	}

	// Merge overlapping or connected ranges
	for i := 0; i < len(ranges)-1; i++ {
		rang := ranges[i]
		for j := i + 1; j < len(ranges); j++ {
			if rang.Overlaps(&ranges[j]) || rang.Connects(&ranges[j]) {
				merged := rang.MergeWith(&ranges[j])
				ranges = util.DeleteAt(ranges, j)
				ranges[i] = merged
				// Repeat the same scan for new range because it now spans further
				i--
				break
			}
		}
	}

	fresh := int64(0)
	for _, rang := range ranges {
		fresh += rang.Length()
	}
	return fresh
}
