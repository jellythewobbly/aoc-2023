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

	p2 := part2(&input)
	fmt.Printf("Part 2: %v\n", p2)
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

func part2(input *[]string) int64 {
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
			sourcesToDestinations = append(sourcesToDestinations, section)
			section = nil
			prevIndex = i
		}
	}

	min := int64(math.MaxInt64)

	for i := 0; i < len(seeds); i += 2 {
		seed, _ := strconv.ParseInt(seeds[i], 10, 64)
		r, _ := strconv.ParseInt(seeds[i+1], 10, 64)

		current := []int64{seed, r}
		for _, section := range sourcesToDestinations {
			current = parseRange(section, current)
		}

		for j := 0; j < len(current); j += 2 {
			if current[j] < min {
				min = current[j]
			}
		}
	}

	return min
}

func parseRange(section [][3]int64, r []int64) []int64 {
	var res []int64

	for i := 0; i < len(r); i += 2 {
		inputStart, inputRange := r[i], r[i+1]
		inputEnd := inputStart + inputRange - 1
		prestine := true
		for _, s := range section {
			sectionStart, sectionRange := s[1], s[2]
			sectionEnd := sectionStart + sectionRange - 1
			diff := s[0] - sectionStart

			if inputStart >= sectionStart && inputEnd <= sectionEnd {
				res = append(res, inputStart+diff, inputRange)
				prestine = false
			} else if inputStart < sectionEnd && inputEnd > sectionEnd {
				if inputStart >= sectionStart {
					res = append(res, inputStart+diff, sectionEnd-inputStart)
				} else {
					res = append(res, sectionStart+diff, sectionRange)
					r = append(r, inputStart, sectionStart-inputStart) // left fall through
				}
				r = append(r, sectionEnd, inputEnd-sectionEnd) // default right fall through
				prestine = false
			} else if inputStart < sectionStart && inputEnd > sectionStart {
				if inputEnd <= sectionEnd {
					res = append(res, sectionStart+diff, inputEnd-sectionStart)
				} else {
					res = append(res, sectionStart+diff, sectionRange)
					r = append(r, sectionEnd, inputEnd-sectionEnd) // right fall through
				}
				r = append(r, inputStart, sectionStart-inputStart) // default left fall through
				prestine = false
			}
		}
		if prestine {
			res = append(res, inputStart, inputRange)
		}
	}

	return res
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
