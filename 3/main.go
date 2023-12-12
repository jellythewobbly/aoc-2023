package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	// "strings"
)

func main() {
	input := getInput()
	p1 := part1(&input)
	fmt.Printf("Part 1: %v\n", p1)

	p2 := part2(&input)
	fmt.Printf("Part 2: %v\n", p2)
}

func part1(input *[]string) int64 {
	sum := int64(0)

	for i, row := range *input {
		start := -1
		adjacent := false
		isDigitMap := make(map[int]bool)

		for j := range row {
			if memoIndexIsDigit(&isDigitMap, &row, j) {
				if start == -1 {
					start = j
				}

				if !adjacent {
					adjacent = isAdjacent(input, i, j)
				}

				if adjacent && (j == len(row)-1 || !memoIndexIsDigit(&isDigitMap, &row, j+1)) {
					num, _ := strconv.ParseInt(row[start:j+1], 10, 64)
					sum += num

					adjacent = false
					start = -1
				}
			} else {
				start = -1
			}
		}
	}
	return sum
}

func part2(input *[]string) int64 {
	sum := int64(0)

	for i, row := range *input {
		for j := range row {
			if row[j] == '*' {
				sum += getGearRatio(input, i, j)
			}
		}
	}

	return sum
}

func getGearRatio(grid *[]string, rowN int, colN int) int64 {
	topCount, midCount, bottomCount := countDistinctNumbers(grid, rowN, colN)
	if topCount+midCount+bottomCount != 2 {
		return int64(0)
	}

	var nums []int64

	if topCount > 0 {
		topNums := horizontalFill((*grid)[rowN-1], colN)
		nums = append(nums, *topNums...)
	}

	if midCount > 0 {
		midNums := horizontalFill((*grid)[rowN], colN)
		nums = append(nums, *midNums...)
	}

	if bottomCount > 0 {
		bottomNums := horizontalFill((*grid)[rowN+1], colN)
		nums = append(nums, *bottomNums...)
	}

	return nums[0] * nums[1]
}

func horizontalFill(row string, colN int) *[]int64 {
	start, end := colN, colN
	var nums []int64
	if isDigit(row[colN]) {
		for start-1 >= 0 && isDigit(row[start-1]) {
			start--
		}

		for end+1 < len(row) && isDigit(row[end+1]) {
			end++
		}
		num, _ := strconv.ParseInt(row[start:end+1], 10, 64)
		return &[]int64{num}
	} else {
		if start-1 >= 0 && isDigit(row[start-1]) {
			start--
			end--
			for start-1 >= 0 && isDigit(row[start-1]) {
				start--
			}
			num, _ := strconv.ParseInt(row[start:end+1], 10, 64)
			nums = append(nums, num)
		}

		start, end := colN, colN
		if end+1 < len(row) && isDigit(row[end+1]) {
			start++
			end++
			for end+1 < len(row) && isDigit(row[end+1]) {
				end++
			}
			num, _ := strconv.ParseInt(row[start:end+1], 10, 64)
			nums = append(nums, num)
		}
	}
	return &nums
}

func countDistinctNumbers(grid *[]string, rowN int, colN int) (int, int, int) {
	topCount := 0
	topPreviousDigit := false
	for _, coordinates := range top {
		r, c := coordinates[0]+rowN, coordinates[1]+colN

		if r < 0 || r == len(*grid) || c < 0 || c == len((*grid)[rowN]) {
			continue
		}

		cell := (*grid)[r][c]
		if isDigit(cell) {
			if !topPreviousDigit {
				topCount++
			}
			topPreviousDigit = true
		} else {
			topPreviousDigit = false
		}
	}

	midCount := 0
	for _, coordinates := range mid {
		r, c := rowN, coordinates[1]+colN

		if c < 0 || c == len((*grid)[rowN]) {
			continue
		}

		cell := (*grid)[r][c]
		if isDigit(cell) {
			midCount++
		}
	}

	bottomCount := 0
	bottomPreviousDigit := false
	for _, coordinates := range bottom {
		r, c := coordinates[0]+rowN, coordinates[1]+colN

		if r < 0 || r == len(*grid) || c < 0 || c == len((*grid)[rowN]) {
			continue
		}

		cell := (*grid)[r][c]
		if isDigit(cell) {
			if !bottomPreviousDigit {
				bottomCount++
			}
			bottomPreviousDigit = true
		} else {
			bottomPreviousDigit = false
		}
	}

	return topCount, midCount, bottomCount
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func memoIndexIsDigit(m *map[int]bool, row *string, index int) bool {
	if (*m)[index] {
		return true
	}
	if digit := isDigit((*row)[index]); digit {
		(*m)[index] = true
		return true
	}
	return false
}

var adjacent = [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
var top = [][]int{{-1, -1}, {-1, 0}, {-1, 1}}
var mid = [][]int{{0, -1}, {0, 1}}
var bottom = [][]int{{1, -1}, {1, 0}, {1, 1}}

func isAdjacent(grid *[]string, row int, col int) bool {
	for _, coordinates := range adjacent {
		r, c := coordinates[0]+row, coordinates[1]+col

		if r < 0 || r == len(*grid) || c < 0 || c == len((*grid)[row]) {
			continue
		}

		cell := (*grid)[r][c]
		if cell != '.' && !isDigit(cell) {
			return true
		}
	}
	return false
}

func getInput() []string {
	fileReader, err := os.Open("./input")
	if err != nil {
		fmt.Println("error in reading file")
		return []string{}
	}

	defer fileReader.Close()

	scanner := bufio.NewScanner(fileReader)

	var res []string

	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	return res
}
