package days

// --- Day 3: Lobby ---

func Day3Part1(lines []string) int {
	sum := 0
	for _, line := range lines {
		maxFirstDigit, maxAtIndex := FindMaxDigitWithin(0, len(line)-1, line)
		maxSecondDigit, _ := FindMaxDigitWithin(maxAtIndex+1, len(line), line)
		sum += maxFirstDigit*10 + maxSecondDigit
	}
	return sum
}

func Day3Part2(lines []string) int64 {
	sum := int64(0)
	for _, line := range lines {
		start := 0
		for d := 12; d >= 1; d-- {
			maxDigit, index := FindMaxDigitWithin(start, len(line)-d+1, line)
			start = index + 1
			sum += int64(maxDigit) * POWERS_OF_10[d-1]
		}
	}
	return sum
}

var POWERS_OF_10 = [13]int64{
	1, 10, 100, 1000, 10_000, 100_000, 1_000_000, 10_000_000, 100_000_000, 1_000_000_000,
	10_000_000_000, 100_000_000_000, 1_000_000_000_000,
}

// FindMaxDigitWithin returns (max, indexOfMax)
// start is inclusive, end exclusive
func FindMaxDigitWithin(start, end int, digits string) (int, int) {
	maxDigit := 0
	indexOfMax := 0
	for i := start; i < end; i++ {
		digit := int(digits[i] - 48)
		if digit > maxDigit {
			maxDigit = digit
			indexOfMax = i
		}
	}
	return maxDigit, indexOfMax
}
