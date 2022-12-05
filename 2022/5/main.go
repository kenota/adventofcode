package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	piles := [][]byte{}

outer:
	for scanner.Scan() {
		l := scanner.Text()

		nCols := (len(l) + 1) / 4

		if len(piles) == 0 {
			piles = make([][]byte, nCols)
		}

		for i := 0; i < nCols; i++ {
			c := l[i*4+1]

			if c >= '0' && c <= '9' {
				// Skip line
				scanner.Scan()
				scanner.Text()
				break outer
			}

			if c != ' ' {
				piles[i] = append(piles[i], c)
			}
		}

	}

	for _, pile := range piles {
		reverse(pile)
	}

	for scanner.Scan() {
		l := scanner.Text()
		var nCount, from, to int
		fmt.Sscanf(l, "move %d from %d to %d", &nCount, &from, &to)

		t := piles[from-1][len(piles[from-1])-nCount:]
		piles[from-1] = piles[from-1][:len(piles[from-1])-nCount]
		// reverse(t)
		piles[to-1] = append(piles[to-1], t...)

		// fmt.Printf("Move %d from %d to %d\n", nCount, from, to)
	}
	for _, pile := range piles {
		fmt.Printf("%c", pile[len(pile)-1])
	}
	println()

}
