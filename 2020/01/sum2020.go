package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func mustReadInput() []int32 {
	var (
		err     error
		res     []int32
		f       *os.File
		scanner *bufio.Scanner
		next    int64
	)

	f, err = os.Open("01-input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		next, err = strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			panic(fmt.Errorf("Failed to parse %s as int: %v", scanner.Text(), err))
		}
		res = append(res, int32(next))
	}

	if err = scanner.Err(); err != nil {
		panic(fmt.Errorf("Failed to read file: %v", err))
	}

	return res

}

func findSumNaive(data []int32, target int32) (*int32, *int32) {
	var (
		i, j        int
		left, right int32
	)

	for i = 0; i < len(data); i++ {
		for j = i + 1; j < len(data); j++ {
			if data[i]+data[j] == target {
				left, right = data[i], data[j]
				return &left, &right
			}
		}
	}

	return nil, nil
}

func findSumMap(data []int32, target int32) (*int32, *int32) {
	// TODO: this will work only if the input is a list of unique numbers
	var (
		m              map[int32]bool
		n, left, right int32
		ok             bool
	)

	m = make(map[int32]bool)

	for _, n = range data {
		m[n] = true
	}

	for _, n = range data {
		right = target - n
		if _, ok = m[target-n]; ok {
			left = n
			return &left, &right
		}
	}

	return nil, nil
}

func findSumThreeNaive(data []int32, target int32) (*int32, *int32, *int32) {
	var (
		first, second, third int32
		i, j, k              int
	)
	for i = 0; i < len(data); i++ {
		for j = i + 1; j < len(data); j++ {
			if data[i]+data[j] > target {
				continue
			}

			for k = j + 1; k < len(data); k++ {
				if data[i]+data[j]+data[k] == target {
					first, second, third = data[i], data[j], data[k]

					return &first, &second, &third
				}
			}
		}
	}

	return nil, nil, nil
}

const target = int32(2020)

func main() {
	var (
		left, right          *int32
		first, second, third *int32
		data                 []int32
	)

	data = mustReadInput()

	left, right = findSumNaive(data, target)
	if left != nil && right != nil {
		fmt.Printf("BruteForce: %d + %d = %d, %d*%d=%d\n", *left, *right, target, *left, *right, (*left)*(*right))
	} else {
		fmt.Printf("No numbers which sum is equal to %d were found", target)
	}

	left, right = findSumMap(data, target)
	if left != nil && right != nil {
		fmt.Printf("HashBased: %d + %d = %d, %d*%d=%d\n", *left, *right, target, *left, *right, (*left)*(*right))
	} else {
		fmt.Printf("No numbers which sum is equal to %d were found", target)
	}

	first, second, third = findSumThreeNaive(data, target)
	if first != nil && second != nil && third != nil {
		fmt.Printf("%d + %d + %d = %d, %d*%d*%d=%d\n",
			*first, *second, *third, target, *first, *second, *third, (*first)*(*second)*(*third))
	}

}
