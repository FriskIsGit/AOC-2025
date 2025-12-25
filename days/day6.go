package days

import (
	"aoc-2025/util"
	"fmt"
	"strings"
)

// --- Day 6: Trash Compactor ---

func Day6Part1(lines []string) int64 {
	var linesOfNums [][]int64
	var signs []byte
	for i, line := range lines {
		if i == len(lines)-1 {
			signs = ParseTerms(line, convertToChar)
			break
		}
		nums := ParseTerms(line, convertToLong)
		linesOfNums = append(linesOfNums, nums)
	}
	if len(signs) != len(linesOfNums[0]) {
		fmt.Println("data is mismatched")
		return -1
	}
	total := int64(0)
	for c := 0; c < len(signs); c++ {
		add := signs[c] == '+'
		var result int64
		if add {
			for row := 0; row < len(linesOfNums); row++ {
				result += linesOfNums[row][c]
			}
		} else {
			result = 1
			for row := 0; row < len(linesOfNums); row++ {
				result *= linesOfNums[row][c]
			}
		}
		total += result
	}
	return total
}

func Day6Part2(lines []string) int64 {
	lineCount := len(lines)
	lineWidth := max(len(lines[0]), len(lines[1]))
	signLine := lines[lineCount-1]

	if lineWidth != len(signLine) {
		fmt.Println("Input data mismatch detected. Extending sign line")
		signLine = extendToLength(signLine, lineWidth)
	}

	total := int64(0)
	for i := lineWidth - 1; i >= 0; i-- {
		if signLine[i] != '+' && signLine[i] != '*' {
			continue
		}
		// Go right until every number ends on whitespace or EOL
		colIndex := findColumnIndexOf(lines, ' ', i) - 1

		// Iterate to the left producing a number from each column
		var nums []int64
		for c := colIndex; c >= i; c-- {
			num := int64(0)
			for row := 0; row < lineCount-1; row++ {
				cell := lines[row][c]
				if cell == ' ' {
					continue
				}
				num *= 10
				digit := cell - 48
				num += int64(digit)
			}
			nums = append(nums, num)
		}

		add := signLine[i] == '+'
		// Add up / Multiply numbers
		result := int64(1)
		if add {
			result = 0
		}
		for _, num := range nums {
			if add {
				result += num
			} else {
				result *= num
			}
		}
		total += result
	}

	return total
}

func extendToLength(str string, length int) string {
	lengthDiff := length - len(str)
	return str + strings.Repeat(" ", lengthDiff)
}

func findColumnIndexOf(grid []string, target byte, from int) int {
	width := len(grid[0])
	for c := from; c < width; c++ {
		allMatch := true
		for _, row := range grid {
			if row[c] != target {
				allMatch = false
				break
			}
		}
		if allMatch {
			return c
		}
	}
	return width
}

func convertToLong(term string) int64 {
	num, _ := util.ParseLong(term)
	return num
}

func convertToChar(term string) byte {
	return term[0]
}

func ParseTerms[T any](line string, convert func(term string) T) []T {
	var terms []T
	contentBuf := strings.Builder{}
	for _, chr := range line {
		if chr == ' ' {
			if contentBuf.Len() > 0 {
				term := convert(contentBuf.String())
				terms = append(terms, term)
				contentBuf.Reset()
			}
			continue
		}
		contentBuf.WriteRune(chr)
	}
	if contentBuf.Len() > 0 {
		term := convert(contentBuf.String())
		terms = append(terms, term)
	}
	return terms
}
