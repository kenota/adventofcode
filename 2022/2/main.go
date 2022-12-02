package main

import (
	"bufio"
	"os"
)

var scoreMatrix [][]int = [][]int{
	{1 + 3, 2 + 6, 3 + 0},
	{1 + 0, 2 + 3, 3 + 6},
	{1 + 6, 2 + 0, 3 + 3},
}

/*
X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win. Good luck!"
*/
var outcomeMatrix [][]int = [][]int{
	{0 + 3, 3 + 1, 6 + 2},
	{0 + 1, 3 + 2, 6 + 3},
	{0 + 2, 3 + 3, 6 + 1},
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	totalScore := 0
	secondScore := 0
	for scanner.Scan() {
		l := scanner.Text()

		// A in hex is 65 and X in hex is 88
		totalScore += scoreMatrix[int(l[0])-65][int(l[2])-88]
		secondScore += outcomeMatrix[int(l[0])-65][int(l[2])-88]
	}

	println(totalScore)
	println(secondScore)

}
