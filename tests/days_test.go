package tests

import (
	"aoc-2025/days"
	"aoc-2025/util"
	"strconv"
	"testing"
)

const INPUTS = "inputs/"
const INPUTS_DEMO = INPUTS + "demo/"

func TestPart1Test(t *testing.T) {
	actual := days.Part1()
	equal(100, actual, t)
}

func loadDayLines(day int) ([]string, error) {
	return util.ReadLines(INPUTS + "day" + strconv.Itoa(day))
}

func equal[T comparable](expected, actual T, t *testing.T) {
	if expected == actual {
		return
	}
	t.Errorf("Expected: %v\nActual:   %v\n", expected, actual)
	t.FailNow()
}
