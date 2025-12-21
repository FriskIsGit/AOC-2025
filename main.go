package main

import (
	"aoc-2025/days"
	"aoc-2025/util"
	"fmt"
	"os"
	"strconv"
)

const INPUTS = "inputs/"
const DEMO = INPUTS + "demo/"

func main() {
	args := os.Args[1:]
	if len(args) <= 1 {
		printHelp()
		return
	}

	day, err := util.ParseInt(args[1])
	if err != nil || day > 12 || day < 1 {
		util.ErrExit("Invalid day number", err)
	}

	switch command := args[0]; command {
	case "run":
		runDay(day)
	case "demo":
	default:
		util.ErrExit("Invalid command", command)
	}
}

func runDay(day int) {
	var resultPart1, resultPart2 any
	switch day {
	case 1:
		lines, err := loadDayLines(day)
		if err != nil {
			util.ErrExit("Missing input", err)
		}
		resultPart1 = days.Day1Part1(lines)
	default:
		util.ErrExit("Unimplemented day:", day)
	}
	fmt.Printf("RESULT (part 1):\n%v", resultPart1)
	fmt.Printf("RESULT (part 2):\n%v", resultPart2)
}

func loadDayLines(day int) ([]string, error) {
	return util.ReadLines(INPUTS + "day" + strconv.Itoa(day) + ".txt")
}

func loadDemoLines(demoName string) ([]string, error) {
	return util.ReadLines(DEMO + demoName)
}

func printHelp() {
	fmt.Println("AOC 2025")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  run [day]")
}
