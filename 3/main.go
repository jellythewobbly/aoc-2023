package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := getInput()
	p1 := part1(&input)
	fmt.Printf("Part 1: %v\n", p1)
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

var adjacent = [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

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
