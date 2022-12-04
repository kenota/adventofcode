package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	includes := 0
	overlaps := 0
	for scanner.Scan() {
		l := scanner.Text()
		var a, b, c, d int
		if _, err := fmt.Sscanf(l, "%d-%d,%d-%d", &a, &b, &c, &d); err != nil {
			panic(err)
		}

		if (d - c) >= (b - a) {
			a, b, c, d = c, d, a, b
		}

		if a <= c && b >= d {
			includes++
		}

		if a > c {
			a, b, c, d = c, d, a, b
		}
		if c <= b {
			overlaps++
		}

	}

	fmt.Printf("Includes: %d\nOverlaps: %d\n", includes, overlaps)
}
