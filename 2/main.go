package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	bag := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	sum := int64(0)

	for i, games := range *input {
		rounds := strings.Split(strings.Split(games, ": ")[1], "; ")
		possible := true
		for _, round := range rounds {
			if !possible {
				break
			}
			sets := strings.Split(round, ", ")
			for _, set := range sets {
				s := strings.Split(set, " ")
				count, _ := strconv.Atoi(s[0])
				colour := s[1]

				if bag[colour] < count {
					possible = false
					break
				}
			}
		}
		if possible {
			sum += int64(i + 1)
		}
	}
	return sum
}

func part2(input *[]string) int64 {
	sum := int64(0)

	for _, games := range *input {
		bag := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		rounds := strings.Split(strings.Split(games, ": ")[1], "; ")
		for _, round := range rounds {
			sets := strings.Split(round, ", ")
			for _, set := range sets {
				s := strings.Split(set, " ")
				count, _ := strconv.Atoi(s[0])
				colour := s[1]
				if bag[colour] < count {
					bag[colour] = count
				}
			}
		}

		num := 1
		for _, count := range bag {
			num *= count
		}
		sum += int64(num)
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
