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
		startNum := GetNearestHalvedNumber(rang.Start)
		startDigitCount := util.NumberOfDigits(startNum)
		startHalfNum := startNum / POWERS_OF_10[startDigitCount/2]

		for halfNum := startHalfNum; ; halfNum++ {
			digitCount := util.NumberOfDigits(halfNum)
			magnitude := POWERS_OF_10[digitCount]
			sequencedNum := halfNum*magnitude + halfNum
			if sequencedNum > rang.End {
				break
			}
			if sequencedNum < rang.Start {
				continue
			}
			invalidSum += sequencedNum
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

var POWERS_OF_10 = [13]int64{
	1, 10, 100, 1000, 10_000, 100_000, 1_000_000, 10_000_000, 100_000_000, 1_000_000_000,
	10_000_000_000, 100_000_000_000, 1_000_000_000_000,
}

func GetNearestHalvedNumber(number int64) int64 {
	digitCount := util.NumberOfDigits(number)
	if digitCount%2 != 0 {
		digitCount++
		number = POWERS_OF_10[digitCount-1]
	}

	magnitudeOfHalf := POWERS_OF_10[digitCount/2]
	leadingDigits := number / magnitudeOfHalf
	return leadingDigits*magnitudeOfHalf + leadingDigits
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
