package main

import (
	"aoc-2025/util"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		printHelp()
		return
	}

	day, err := util.ParseInt(args[1])
	if err != nil || day > 12 || day < 1 {
		util.ErrExit("Invalid day number", err)
	}

	switch command := args[0]; command {
	case "run":
	case "demo":
	default:
		util.ErrExit("Invalid command", command)
	}
}

func printHelp() {
	fmt.Println("AOC 2025")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  run [day]")
	fmt.Println("  demo [day]")
}
