package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadLines(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func ReadFileToString(path string) (string, error) {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(fileBytes), nil
}

func ParseInt(number string) (int, error) {
	num, err := strconv.ParseInt(number, 10, 32)
	if err != nil {
		return 0, err
	}
	return int(num), nil
}

func ErrExit(messages ...any) {
	fmt.Println(messages)
	os.Exit(1)
}
