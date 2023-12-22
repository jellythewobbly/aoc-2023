package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
	regex := regexp.MustCompile(" +")
	time := regex.Split((*input)[0], -1)[1:]
	distance := regex.Split((*input)[1], -1)[1:]

	res := int64(1)

	for i := 0; i < len(time); i++ {
		t, _ := strconv.ParseInt(time[i], 10, 64)
		d, _ := strconv.ParseInt(distance[i], 10, 64)

		isExact, charge := bSearch(d, t)
		ways := t + 1 - 2*charge
		if isExact {
			ways -= 2
		}
		res *= ways
	}

	return res
}

func part2(input *[]string) int64 {
	regex := regexp.MustCompile(" +")
	time := strings.Join(regex.Split((*input)[0], -1)[1:], "")
	distance := strings.Join(regex.Split((*input)[1], -1)[1:], "")

	t, _ := strconv.ParseInt(time, 10, 64)
	d, _ := strconv.ParseInt(distance, 10, 64)

	isExact, charge := bSearch(d, t)
	ways := t + 1 - 2*charge
	if isExact {
		ways -= 2
	}

	return ways
}

func bSearch(d int64, t int64) (isExact bool, charge int64) {
	low, high := int64(0), t
	for low <= high {
		mid := low + (high-low)/2
		distance := calculateDistance(t, mid)
		if distance == d {
			return true, mid
		}
		if distance < d {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false, low
}

func calculateDistance(total int64, charge int64) int64 {
	return charge * (total - charge)
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
