package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
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

	for _, row := range *input {
		splits := strings.Split(row, " | ")
		winners := strings.Split(strings.Split(splits[0], ": ")[1], " ")
		entries := strings.Split(splits[1], " ")
		winMap := make(map[string]bool)

		for _, winner := range winners {
			winMap[winner] = true
		}

		count := 0

		for _, entry := range entries {
			if entry != "" {
				if _, ok := winMap[entry]; ok {
					count++
				}
			}
		}

		sum += int64(math.Pow(2, float64(count-1)))
	}

	return sum
}

func part2(input *[]string) int64 {
	cardCollection := make([]int, len(*input))

	for i, row := range *input {
		splits := strings.Split(row, " | ")
		winners := strings.Split(strings.Split(splits[0], ": ")[1], " ")
		entries := strings.Split(splits[1], " ")
		winMap := make(map[string]bool)

		for _, winner := range winners {
			winMap[winner] = true
		}

		count := 0

		for _, entry := range entries {
			if entry != "" {
				if _, ok := winMap[entry]; ok {
					count++
				}
			}
		}

		wonCardNumber := i + 1
		currentCardCount := cardCollection[i] + 1
		for count > 0 {
			cardCollection[wonCardNumber] += currentCardCount
			wonCardNumber++
			count--
		}
	}

	sum := int64(len(*input))
	for _, val := range cardCollection {
		sum += int64(val)
	}

	return sum
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
