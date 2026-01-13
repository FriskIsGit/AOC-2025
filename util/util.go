package util

import (
	"bufio"
	"cmp"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
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

func ToBoard(lines []string) [][]byte {
	rowCount := len(lines)
	width := len(lines[0])
	board := make([][]byte, rowCount)
	for r := 0; r < rowCount; r++ {
		row := make([]byte, width)
		copy(row, lines[r])
		board[r] = row
	}
	return board
}

func PrintBoard(board [][]byte) {
	for _, row := range board {
		fmt.Println(string(row))
	}
}

func PrintIntBoard(board [][]int) {
	for _, row := range board {
		fmt.Println(row)
	}
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

type Range struct {
	Start int64
	End   int64
}

func NewRange(start, end int64) Range {
	return Range{start, end}
}

func ParseRange(strRange string) Range {
	left, right, _ := strings.Cut(strRange, "-")
	start, _ := ParseLong(left)
	end, _ := ParseLong(right)
	return NewRange(start, end)
}

func (r *Range) MergeWith(other *Range) Range {
	mergedStart := min(r.Start, other.Start)
	mergedEnd := max(r.End, other.End)
	return Range{Start: mergedStart, End: mergedEnd}
}

func (r *Range) Overlaps(other *Range) bool {
	return r.Start <= other.End && other.Start <= r.End
}

func (r *Range) Connects(other *Range) bool {
	return r.End+1 == other.Start || r.Start == other.End+1
}

func (r *Range) Includes(value int64) bool {
	return r.Start <= value && value <= r.End
}

func (r *Range) Length() int64 {
	return r.End - r.Start + 1
}

func DeleteAt[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func LargestN[T cmp.Ordered](arr []T, n int) []T {
	if n == 0 {
		return []T{}
	}
	maxes := make([]T, n, n+1)
	for _, t := range arr {
		for index, high := range maxes {
			if t >= high {
				slices.Insert(maxes, index, t)
				maxes = maxes[:n]
				break
			}
		}
	}
	return maxes
}

func Distance2D(x, y, x1, y1 int) float64 {
	xDelta := float64(x - x1)
	yDelta := float64(y - y1)
	return math.Sqrt(xDelta*xDelta + yDelta*yDelta)
}

type Set[T comparable] struct {
	Set map[T]struct{}
}

func (s *Set[T]) Contains(el T) bool {
	_, contains := s.Set[el]
	return contains
}

func (s *Set[T]) Add(el T) {
	s.Set[el] = struct{}{}
}

func (s *Set[T]) Remove(el T) {
	delete(s.Set, el)
}

func (s *Set[T]) Size() int {
	return len(s.Set)
}

func NewSet[T comparable](capacity int) *Set[T] {
	return &Set[T]{
		make(map[T]struct{}, capacity),
	}
}

func (s *Set[T]) AddAll(other *Set[T]) {
	if s == nil || other == nil {
		return
	}
	for k := range other.Set {
		s.Add(k)
	}
}

func TrimLeftAndRight(s string, left, right string) string {
	s = strings.TrimLeft(s, left)
	return strings.TrimRight(s, right)
}

func MemZeroBoolArray(b []bool) {
	for i := range b {
		b[i] = false
	}
}

func CopyAppend[T any](source []T, newEl T) []T {
	dest := make([]T, len(source)+1)
	copy(dest, source)
	dest[len(dest)-1] = newEl
	return dest
}

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | uint | uint8 | uint16 | uint32 | uint64
}

func HashSlice[T Number](slice []T) T {
	hash := T(0)
	prime := T(31) // Choose a prime number for better distribution

	for _, value := range slice {
		hash = hash*prime + value
	}
	return hash
}

var POWERS_OF_2 = [18]int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536, 131072}

func AllCombinations[T any](elements []T) [][]T {
	allCombinationsCount := POWERS_OF_2[len(elements)]
	combinations := make([][]T, 0, allCombinationsCount)
	nextCombinations := make([][]T, 0, allCombinationsCount)

	combinations = append(combinations, []T{})
	for _, el := range elements {
		for _, comb := range combinations {
			newComb := CopyAppend(comb, el)
			nextCombinations = append(nextCombinations, comb, newComb)
		}
		combinations, nextCombinations = nextCombinations, combinations
		nextCombinations = nextCombinations[:0]
	}
	return combinations
}
