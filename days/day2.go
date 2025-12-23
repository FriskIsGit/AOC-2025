package days

import (
	"aoc-2025/util"
	"strings"
)

// --- Day 2: Gift Shop ---

func Day2Part1(input string) int64 {
	ranges := convertToRanges(input)

	invalidSum := int64(0)
	for _, rang := range ranges {
		for val := rang.Start; val <= rang.End; val++ {
			strValue := util.LongToString(val)
			digitCount := len(strValue)
			if digitCount%2 != 0 {
				continue
			}
			half := digitCount / 2
			leftNum, rightNum := strValue[:half], strValue[half:]
			if leftNum == rightNum {
				invalidSum += val
			}
		}
	}
	return invalidSum
}

func Day2Part2(input string) int64 {
	ranges := convertToRanges(input)

	invalidSum := int64(0)
	for _, rang := range ranges {
		for val := rang.Start; val <= rang.End; val++ {
			strValue := util.LongToString(val)
			digitCount := len(strValue)
			half := digitCount / 2
			for length := 1; length <= half; length++ {
				pattern := strValue[:length]
				if IsSequenceOf(strValue, pattern) {
					invalidSum += val
					break
				}
			}
		}
	}
	return invalidSum
}

// IsSequenceOf checks if str is a sequence of pattern
func IsSequenceOf(str, pattern string) bool {
	jump := len(pattern)
	if len(str)%jump != 0 {
		return false
	}
	repeats := len(str) / jump
	offset := 0

	for i := 0; i < repeats; i++ {
		if str[offset:offset+jump] != pattern {
			return false
		}
		offset += jump
	}
	return true
}

func convertToRanges(input string) []util.Range {
	strRanges := strings.Split(input, ",")
	var ranges []util.Range
	for _, strRange := range strRanges {
		r := util.ParseRange(strRange)
		ranges = append(ranges, r)
	}
	return ranges
}
