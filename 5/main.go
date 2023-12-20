package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := getInput()
	p1 := part1(&input)
	fmt.Printf("Part 1: %v\n", p1)
}

func part1(input *[]string) int64 {
	seeds := strings.Split(strings.Split((*input)[0], ": ")[1], " ")
	var sourcesToDestinations [][][3]int64

	*input = append(*input, "")
	prevIndex := 1
	rows := len(*input)
	for i := 2; i < rows; i++ {
		row := (*input)[i]
		if row == "" {
			j := prevIndex + 2
			var section [][3]int64
			for j < i {
				split := strings.Split((*input)[j], " ")
				d, _ := strconv.ParseInt(split[0], 10, 64)
				s, _ := strconv.ParseInt(split[1], 10, 64)
				r, _ := strconv.ParseInt(split[2], 10, 64)
				section = append(section, [3]int64{d, s, r})
				j++
			}
			sort.Slice(section, func(i, j int) bool {
				return section[i][1] < section[j][1]
			})
			sourcesToDestinations = append(sourcesToDestinations, section)
			section = nil
			prevIndex = i
		}
	}

	min := int64(math.MaxInt64)

	for _, seed := range seeds {
		current, _ := strconv.ParseInt(seed, 10, 64)
		for _, section := range sourcesToDestinations {
			current = bSearch(section, current)
		}
		if current < min {
			min = current
		}
	}

	return min
}

func bSearch(arr [][3]int64, val int64) int64 {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)/2
		start := arr[mid][1]
		end := start + arr[mid][2] - 1
		if val >= start && val <= end {
			diff := arr[mid][0] - arr[mid][1]
			return val + diff
		}
		if val < start {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return val
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
