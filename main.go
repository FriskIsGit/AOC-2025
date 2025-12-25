package main

import (
	"aoc-2025/days"
	"aoc-2025/util"
	"fmt"
	"os"
	"strconv"
	"time"
)

const INPUTS = "inputs/"
const DEMO = INPUTS + "demo/"
const DAYS = "days/"

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
	case "create":
		createDayIfMissing(day)
	default:
		util.ErrExit("Invalid command", command)
	}
}

func createDayIfMissing(day int) {
	if err := os.MkdirAll(DEMO, os.ModePerm); err != nil {
		util.ErrExit("Failed to create output directory", err)
	}
	strDay := strconv.Itoa(day)
	inputFileName := "day" + strDay + ".txt"
	util.CreateEmptyFile(INPUTS + inputFileName)
	util.CreateEmptyFile(DEMO + inputFileName)
	util.CreateEmptyFile(DAYS + "day" + strDay + ".go")
}

func runDay(day int) {
	var resultPart1, resultPart2 any
	lines, err := loadDayLines(day)
	if err != nil {
		util.ErrExit("Missing input", err)
	}

	start := time.Now()
	switch day {
	case 1:
		resultPart1 = days.Day1Part1(lines)
		resultPart2 = days.Day1Part2(lines)
	case 2:
		input := lines[0]
		resultPart1 = days.Day2Part1(input)
		part2Start := time.Now()
		resultPart2 = days.Day2Part2(input)
		part2Taken := time.Since(part2Start)
		fmt.Printf("\nTime taken day2 part2: %v", part2Taken)
	case 3:
		resultPart1 = days.Day3Part1(lines)
		resultPart2 = days.Day3Part2(lines)
	case 4:
		resultPart1 = days.Day4Part1(lines)
		resultPart2 = days.Day4Part2(lines)
	case 5:
		resultPart1 = days.Day5Part1(lines)
		resultPart2 = days.Day5Part2(lines)
	case 6:
		resultPart1 = days.Day6Part1(lines)
		resultPart2 = days.Day6Part2(lines)
	default:
		util.ErrExit("Unimplemented day:", day)
	}
	timeTaken := time.Since(start)
	fmt.Printf("\nTime taken: %v\n", timeTaken)
	fmt.Printf("RESULT (part 1): %v\n", resultPart1)
	fmt.Printf("RESULT (part 2): %v\n", resultPart2)
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
	fmt.Println("  create [day]")
}
