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
	equal(966, actual, t)
	// 966 is too low
}

func TestDay1Part1Demo(t *testing.T) {
	lines, err := loadDemoLines("day1p1.txt")
	if err != nil {
		t.Error(err)
		return
	}
	actual := days.Day1Part1(lines)
	equal(3, actual, t)
}

func TestDay1Part1Custom(t *testing.T) {
	lines, err := loadCustomLines("custom1p1.txt")
	if err != nil {
		t.Error(err)
		return
	}
	actual := days.Day1Part1(lines)
	equal(1, actual, t)
}

func TestNextDialPos(t *testing.T) {
	nextPos := days.NextDialPos(0, -150)
	equal(50, nextPos, t)
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
