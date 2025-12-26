package tests

import (
	"aoc-2025/days"
	"aoc-2025/util"
	"strconv"
	"testing"
)

const INPUTS = "../inputs/"
const DEMO = INPUTS + "demo/"
const CUSTOM = INPUTS + "custom/"

// ----- DAY 1 -----

func TestDay1Part1Full(t *testing.T) {
	lines, err := loadDayLines(1)
	if err != nil {
		t.Error(err)
		return
	}
	actual := days.Day1Part1(lines)
	equal(1078, actual, t)
}

func TestDay1Part2Full(t *testing.T) {
	lines, err := loadDayLines(1)
	if err != nil {
		t.Error(err)
		return
	}
	actual := days.Day1Part2(lines)
	equal(6412, actual, t)
}

func TestDay1Part1Demo(t *testing.T) {
	lines, err := loadDemoLines("day1.txt")
	if err != nil {
		t.Error(err)
		return
	}
	actual := days.Day1Part1(lines)
	equal(3, actual, t)
}

func TestDay1Part2Demo(t *testing.T) {
	lines, err := loadDemoLines("day1.txt")
	if err != nil {
		t.Error(err)
		return
	}
	actual := days.Day1Part2(lines)
	equal(6, actual, t)
}

func TestGetZeroHitsExpect10(t *testing.T) {
	zeroHits := days.GetZeroHits(50, 1000)
	equal(10, zeroHits, t)
}

func TestGetZeroHitsNegative(t *testing.T) {
	zeroHits := days.GetZeroHits(50, -250)
	equal(3, zeroHits, t)
}

func TestGetZeroHitsFrom0To200(t *testing.T) {
	zeroHits := days.GetZeroHits(0, 200)
	equal(2, zeroHits, t)
}

func TestGetZeroHitsFrom0ToNegative200(t *testing.T) {
	zeroHits := days.GetZeroHits(0, -200)
	equal(2, zeroHits, t)
}

func TestGetZeroHitsExpect1(t *testing.T) {
	zeroHits := days.GetZeroHits(20, -40)
	equal(1, zeroHits, t)
}

func TestGetZeroHitsCustom(t *testing.T) {
	zeroHits := days.GetZeroHits(50, -51)
	equal(1, zeroHits, t)
}

func TestGetZeroHitsMoveFromZero(t *testing.T) {
	zeroHits := days.GetZeroHits(0, -5)
	equal(0, zeroHits, t)
}

// ----- DAY 2 -----

func TestDay2Part1Full(t *testing.T) {
	input, err := loadDayToString(2)
	if err != nil {
		t.Error(err)
		return
	}
	actual := days.Day2Part1(input)
	equal(37314786486, actual, t)
}

func TestDay2Part2Full(t *testing.T) {
	input, err := loadDayToString(2)
	if err != nil {
		t.Error(err)
		return
	}
	actual := days.Day2Part2(input)
	equal(47477053982, actual, t)
}

func TestDay2Part1Demo(t *testing.T) {
	input, err := loadDemoToString("day2.txt")
	if err != nil {
		t.Error(err)
		return
	}
	actual := days.Day2Part1(input)
	equal(1227775554, actual, t)
}

func TestDay2Part2Demo(t *testing.T) {
	input, err := loadDemoToString("day2.txt")
	if err != nil {
		t.Error(err)
		return
	}
	actual := days.Day2Part2(input)
	equal(4174379265, actual, t)
}

func TestIsSequenceOfExpectTrue(t *testing.T) {
	result := days.IsSequenceOf("abcabc", "abc")
	equalsTrue(result, t)
}

func TestIsSequenceOfExpectFalse(t *testing.T) {
	result := days.IsSequenceOf("abcabcxyz", "abc")
	equalsFalse(result, t)
}

// ----- DAY 3 -----

func TestDay3Part1Full(t *testing.T) {
	lines, err := loadDayLines(3)
	if err != nil {
		t.Error(err)
		return
	}
	actual := days.Day3Part1(lines)
	equal(17278, actual, t)
}

func TestDay3Part2Full(t *testing.T) {
	lines, err := loadDayLines(3)
	if err != nil {
		t.Error(err)
		return
	}
	actual := days.Day3Part2(lines)
	equal(171528556468625, actual, t)
}

func TestDay3Part1Demo(t *testing.T) {
	lines, err := loadDemoLines("day3.txt")
	if err != nil {
		t.Error(err)
		return
	}
	actual := days.Day3Part1(lines)
	equal(357, actual, t)
}

func TestDay3Part2Demo(t *testing.T) {
	lines, err := loadDemoLines("day3.txt")
	if err != nil {
		t.Error(err)
		return
	}
	actual := days.Day3Part2(lines)
	equal(3121910778619, actual, t)
}

// ----- DAY 4 -----

func TestDay4Part1Full(t *testing.T) {
	lines, err := loadDayLines(4)
	if err != nil {
		t.Error(err)
		return
	}
	actual := days.Day4Part1(lines)
	equal(1441, actual, t)
}

func TestDay4Part1Demo(t *testing.T) {
	lines, err := loadDemoLines("day4.txt")
	if err != nil {
		t.Error(err)
		return
	}
	actual := days.Day4Part1(lines)
	equal(13, actual, t)
}

func TestDay4Part2Full(t *testing.T) {
	lines, err := loadDayLines(4)
	if err != nil {
		t.Error(err)
		return
	}
	actual := days.Day4Part2(lines)
	equal(9050, actual, t)
}

func TestDay4Part2Demo(t *testing.T) {
	lines, err := loadDemoLines("day4.txt")
	if err != nil {
		t.Error(err)
		return
	}
	actual := days.Day4Part2(lines)
	equal(43, actual, t)
}

// ----- DAY 5 -----

func TestDay5Part1Full(t *testing.T) {
	lines, err := loadDayLines(5)
	if err != nil {
		t.Error(err)
		return
	}
	actual := days.Day5Part1(lines)
	equal(513, actual, t)
}

func TestDay5Part1Demo(t *testing.T) {
	lines, err := loadDemoLines("day5.txt")
	if err != nil {
		t.Error(err)
		return
	}
	actual := days.Day5Part1(lines)
	equal(3, actual, t)
}

func TestDay5Part2Full(t *testing.T) {
	lines, err := loadDayLines(5)
	if err != nil {
		t.Error(err)
		return
	}
	actual := days.Day5Part2(lines)
	equal(339668510830757, actual, t)
}

func TestDay5Part2Demo(t *testing.T) {
	lines, err := loadDemoLines("day5.txt")
	if err != nil {
		t.Error(err)
		return
	}
	actual := days.Day5Part2(lines)
	equal(14, actual, t)
}

// ----- DAY 6 -----

func TestDay6Part1Full(t *testing.T) {
	lines, err := loadDayLines(6)
	if err != nil {
		t.Error(err)
		return
	}
	actual := days.Day6Part1(lines)
	equal(4580995422905, actual, t)
}

func TestDay6Part1Demo(t *testing.T) {
	lines, err := loadDemoLines("day6.txt")
	if err != nil {
		t.Error(err)
		return
	}
	actual := days.Day6Part1(lines)
	equal(4277556, actual, t)
}

func TestDay6Part2Full(t *testing.T) {
	lines, err := loadDayLines(6)
	if err != nil {
		t.Error(err)
		return
	}
	actual := days.Day6Part2(lines)
	equal(10875057285868, actual, t)
}

func TestDay6Part2Demo(t *testing.T) {
	lines, err := loadDemoLines("day6.txt")
	if err != nil {
		t.Error(err)
		return
	}
	actual := days.Day6Part2(lines)
	equal(3263827, actual, t)
}

// ----- DAY 7 -----

func TestDay7Part1Demo(t *testing.T) {
	lines, err := loadDemoLines("day7.txt")
	if err != nil {
		t.Error(err)
		return
	}
	actual := days.Day7Part1(lines)
	equal(21, actual, t)
}

// --- Util functions ---

func loadDayLines(day int) ([]string, error) {
	return util.ReadLines(INPUTS + "day" + strconv.Itoa(day) + ".txt")
}

func loadDayToString(day int) (string, error) {
	return util.ReadFileToString(INPUTS + "day" + strconv.Itoa(day) + ".txt")
}

func loadDemoToString(demoName string) (string, error) {
	return util.ReadFileToString(DEMO + demoName)
}

func loadDemoLines(demoName string) ([]string, error) {
	return util.ReadLines(DEMO + demoName)
}

func loadCustomLines(customName string) ([]string, error) {
	return util.ReadLines(CUSTOM + customName)
}

func equal[T comparable](expected, actual T, t *testing.T) {
	if expected == actual {
		return
	}
	t.Errorf("\nExpected: %v\nActual:   %v\n", expected, actual)
	t.FailNow()
}

func equalsTrue(actual bool, t *testing.T) {
	if actual {
		return
	}
	t.Errorf("\nExpected: %v\nActual:   %v\n", true, actual)
	t.FailNow()
}

func equalsFalse(actual bool, t *testing.T) {
	if !actual {
		return
	}
	t.Errorf("\nExpected: %v\nActual:   %v\n", false, actual)
	t.FailNow()
}
