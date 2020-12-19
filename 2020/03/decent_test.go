package main

import (
	"log"
	"strings"
	"testing"
)

const field = `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#
`

func TestTaskField(t *testing.T) {
	const expected = 7
	field := mustReadField(strings.NewReader(field))

	paths := findPath(field)
	if paths != expected {
		log.Printf("Expected to find %d trees, found %d", expected, paths)
		t.FailNow()
	}
}
