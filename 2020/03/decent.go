package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func mustReadField(r io.Reader) [][]byte {
	var (
		res     [][]byte
		err     error
		scanner *bufio.Scanner
		line    string
		i       int
		next    []byte
	)

	scanner = bufio.NewScanner(r)
	for scanner.Scan() {
		next = make([]byte, 0)

		line = scanner.Text()
		for i = 0; i < len(line); i++ {
			if line[i] == '#' {
				next = append(next, 1)
			} else if line[i] == '.' {
				next = append(next, 0)
			} else {
				log.Fatalf("Uknown character %v at pos %d in %s", line[i], i, line)
			}
		}
		res = append(res, next)
	}

	if err = scanner.Err(); err != nil {
		log.Fatalf("error reading field: %v", err)
	}

	return res
}

func findPath(field [][]byte, right, down int) int {
	var (
		trees, x, y int
	)

	y += down
	for y < len(field) {
		x += right
		if x >= len(field[y]) {
			x = x - len(field[y])
		}
		if field[y][x] == 1 {
			trees++
		}
		y += down
	}

	return trees
}

func main() {
	var (
		f       *os.File
		err     error
		field   [][]byte
		trees11 int
		trees31 int
		trees51 int
		trees71 int
		trees12 int
	)

	f, err = os.Open("03-input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	field = mustReadField(f)

	trees11 = findPath(field, 1, 1)
	fmt.Printf("Number of trees struck via 1+1 route: %d\n", trees11)

	trees31 = findPath(field, 3, 1)
	fmt.Printf("Number of trees struck via 3+1 route (default): %d\n", trees31)

	trees51 = findPath(field, 5, 1)
	fmt.Printf("Number of trees struck via 5+1 path: %d\n", trees51)

	trees71 = findPath(field, 7, 1)
	fmt.Printf("Number of trees struck via 7+1 path: %d\n", trees71)

	trees12 = findPath(field, 1, 2)
	fmt.Printf("Number of trees struck on 1+2 path: %d\n", trees12)

	// incorrect: 1277702668242%
	fmt.Printf("Total trees encountered on all paths: %d\n", trees11*trees31*trees51*trees71*trees12)

}
