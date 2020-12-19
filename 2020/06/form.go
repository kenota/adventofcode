package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	var (
		f                  *os.File
		err                error
		data               []byte
		group              string
		totalAnswers       int
		totalCommonAnswers int
	)

	if f, err = os.Open("06-input.txt"); err != nil {
		panic(err)
	}
	defer f.Close()

	if data, err = ioutil.ReadAll(f); err != nil {
		panic(err)
	}

	for _, group = range strings.Split(string(data), "\n\n") {
		var answers = make(map[rune]bool)
		for _, c := range group {
			if c == '\n' {
				continue
			}
			answers[c] = true
		}
		totalAnswers += len(answers)
	}

	fmt.Printf("Total answers: %d\n", totalAnswers)

	for _, group = range strings.Split(string(data), "\n\n") {
		var groupCommonAnswers map[rune]bool
		for _, person := range strings.Split(group, "\n") {
			var answers = make(map[rune]bool)
			for _, c := range person {
				answers[c] = true
			}
			if groupCommonAnswers == nil {
				groupCommonAnswers = answers
			} else {
				groupCommonAnswers = setAnd(groupCommonAnswers, answers)
			}
		}
		totalCommonAnswers += len(groupCommonAnswers)
	}

	fmt.Printf("Total common yes answers: %d\n", totalCommonAnswers)

}

func setAnd(left, right map[rune]bool) map[rune]bool {
	var (
		res = make(map[rune]bool)
		k   rune
		ok  bool
	)

	for k = range left {
		if _, ok = right[k]; ok {
			res[k] = true
		}
	}

	return res
}
