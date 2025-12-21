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
	equal(5337, actual, t)
	// 5337 is too low, 6955 is too high
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

func TestDay2Part1Full(t *testing.T) {
	input, err := loadDayToString(1)
	if err != nil {
		t.Error(err)
		return
	}
	actual := days.Day2Part1(input)
	equal(-1, actual, t)
}

func loadDayLines(day int) ([]string, error) {
	return util.ReadLines(INPUTS + "day" + strconv.Itoa(day) + ".txt")
}

func loadDayToString(day int) (string, error) {
	return util.ReadFileToString(INPUTS + "day" + strconv.Itoa(day) + ".txt")
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
	t.Errorf("Expected: %v\nActual:   %v\n", expected, actual)
	t.FailNow()
}
