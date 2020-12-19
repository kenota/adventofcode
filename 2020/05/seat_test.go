package main

import (
	"fmt"
	"log"
	"testing"
)

func assert(cond bool, msg string, t *testing.T) {
	if cond == false {
		log.Printf("Assert failed: %s", msg)
		t.FailNow()
	}
}

func TestFindingRow(t *testing.T) {
	var row int

	row = binarySearch(0, 127, "FBFBBFF")
	assert(row == 44, fmt.Sprintf("Expecting row to be 44, got %d", row), t)
}

func TestFindingSeat(t *testing.T) {
	var seat int

	seat = binarySearch(0, 7, "RLR")
	assert(seat == 5, fmt.Sprintf("Expecting seat to be 5, got %d", seat), t)
}

func TestFindSeatID(t *testing.T) {
	var testCases = map[string]int{
		"BFFFBBFRRR": 567,
		"FFFBBBFRRR": 119,
		"BBFFBBFRLL": 820,
		"FBFBBFFRLR": 357,
	}

	for k, v := range testCases {
		assert(findSeatID(k) == v, fmt.Sprintf("Expecting seat id %d for seat '%s', got: %d instead", v, k, findSeatID(k)), t)
	}

}
