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

	p2 := part2(&input)
	fmt.Printf("Part 2: %v\n", p2)
}

func part1(input *[]string) int64 {
	sum := int64(0)
	for _, line := range *input {
		first, last := int64(-1), int64(-1)
		for i := range line {
			if first == -1 {
				start := line[i]
				if ok, value := parseByteInt64(&start); ok {
					first = value
				}
			}

			if last == -1 {
				end := line[len(line)-1-i]
				if ok, value := parseByteInt64(&end); ok {
					last = value
				}
			}

			if first != -1 && last != -1 {
				sum += (first*10 + last)
				break
			}
		}
	}
	return sum
}

func part2(input *[]string) int64 {
	sum := int64(0)
	for _, line := range *input {
		first, last := int64(-1), int64(-1)
		for i := range line {
			if first == -1 {
				start := line[i]
				if ok, value := parseByteInt64(&start); ok {
					first = value
				} else if ok, value := parseStringInt64(&line, i, DIRECTION_FORWARD); ok {
					first = value
				}
			}

			if last == -1 {
				endIndex := len(line) - 1 - i
				end := line[endIndex]
				if ok, value := parseByteInt64(&end); ok {
					last = value
				} else if ok, value := parseStringInt64(&line, endIndex, DIRECTION_BACKWARD); ok {
					last = value
				}
			}

			if first != -1 && last != -1 {
				sum += (first*10 + last)
				break
			}
		}
	}
	return sum
}

func parseByteInt64(b *byte) (success bool, value int64) {
	if *b >= '0' && *b <= '9' {
		i64, _ := strconv.ParseInt(string(*b), 10, 64)
		return true, i64
	}
	return false, 0
}

type Direction int8

const (
	DIRECTION_FORWARD Direction = iota
	DIRECTION_BACKWARD
)

func parseStringInt64(s *string, index int, direction Direction) (success bool, value int64) {
	digitWords := map[string]int64{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for _, v := range []int{3, 4, 5} {
		startIndex := index
		endIndex := index + v

		if direction == DIRECTION_BACKWARD {
			offset := 1
			startIndex = index - v + offset
			endIndex = index + offset
		}

		if startIndex < 0 || endIndex > len(*s) {
			return false, 0
		}
		sliced := (*s)[startIndex:endIndex]

		if val, ok := digitWords[sliced]; ok {
			return true, val
		}
	}
	return false, 0
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
