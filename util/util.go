package util

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func NumberOfDigits(posNum int64) int {
	if posNum == 0 {
		return 1
	}
	log10 := math.Log10(float64(posNum))
	return int(math.Floor(log10)) + 1
}

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

func ParseLong(number string) (int64, error) {
	num, err := strconv.ParseInt(number, 10, 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func LongToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

func ErrExit(messages ...any) {
	fmt.Println(messages)
	os.Exit(1)
}

func CreateEmptyFile(path string) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_EXCL, os.ModePerm)
	if err == nil {
		f.Close()
	}
}
