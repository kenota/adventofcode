package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
	We need to find an intersection of a set. Just going by map should be fine
*/

func findCommon(input []string) int {
	var accum []byte
	for _, l := range input {
		arr := make([]byte, 255)
		for _, c := range l {
			arr[int(c)] = 1
		}
		if accum == nil {
			accum = arr
		} else {
			for i := range accum {
				accum[i] = accum[i] & arr[i]
			}
		}
	}

	for i, v := range accum {
		if v > 0 {
			if i <= 'Z' {
				return int(i) - 'A' + 27
			} else {
				return int(i) - 'a' + 1
			}
		}
	}

	panic("no common")
}

func main() {
	r := bufio.NewScanner(os.Stdin)

	total := 0
	badgeTotal := 0

	lines := []string{}
	for r.Scan() {
		l := r.Text()

		total += findCommon([]string{l[0 : len(l)/2], l[len(l)/2:]})
		lines = append(lines, l)
		if len(lines) == 3 {
			badgeTotal += findCommon(lines)
			lines = lines[:0]
		}
	}

	fmt.Printf("Rucksack priority priority: %d\nBadges prioority: %d\n", total, badgeTotal)
}
