package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func binarySearch(low, high int, path string) int {
	var (
		c      rune
		middle int
	)

	for _, c = range path {
		middle = (high - low) / 2
		switch c {
		case 'B', 'R':
			// upper half
			if high-low == 1 {
				return high
			}
			low = low + middle + 1

		case 'F', 'L':
			// lower half
			if high-low == 1 {
				return low
			}
			high = high - middle - 1
		default:
			panic(fmt.Sprintf("Uknown char: %v", c))
		}
	}
	panic(fmt.Sprintf("Not full path, low: %d high: %d", low, high))
}

func findSeatID(s string) int {
	if len(s) != 10 {
		panic(fmt.Sprintf("Incorrect seat spec, expected 10 charatcters. Got: '%s'", s))
	}

	return binarySearch(0, 127, s[:7])*8 + binarySearch(0, 7, s[7:])
}

func main() {
	var (
		f         *os.File
		err       error
		scanner   *bufio.Scanner
		id, maxid int
		seats     []int
		i         int
	)

	f, err = os.Open("05-input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		if id = findSeatID(scanner.Text()); id > maxid {
			maxid = id
		}
		seats = append(seats, id)
	}

	if err = scanner.Err(); err != nil {
		panic(fmt.Sprintf("Failed to read input data: %v", err))
	}

	fmt.Printf("Max seat id: %d\n", maxid)

	sort.Ints(seats)
	for i = 1; i < len(seats); i++ {
		if seats[i]-seats[i-1] > 1 {
			fmt.Printf("My seat is %d\n", seats[i]-1)
		}
	}

}
