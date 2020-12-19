package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type passwordPolicy struct {
	min      int
	max      int
	char     byte
	password string
}

func parsePasswordPolicy(s string) (*passwordPolicy, error) {
	var (
		res passwordPolicy
		err error
	)
	if _, err = fmt.Fscanf(strings.NewReader(s), "%d-%d %c: %s", &res.min, &res.max, &res.char, &res.password); err != nil {
		return nil, fmt.Errorf("failed to scan input: %s:  %v", s, err)
	}

	if res.min > res.max {
		return nil, fmt.Errorf("Min is not less than max values min:max are: %d:%d", res.min, res.max)
	}

	return &res, nil
}

func (p *passwordPolicy) isValid() bool {
	var (
		i, count int
	)

	for i = 0; i < len(p.password); i++ {
		if p.password[i] == p.char {
			count++
			if count > p.max {
				return false
			}
		}
	}

	return count >= p.min
}

func (p *passwordPolicy) isValidV2() bool {
	var (
		firstIdx, secondIdx int
	)
	firstIdx, secondIdx = p.min-1, p.max-1

	if firstIdx > len(p.password) {
		return false
	}

	if p.password[firstIdx] == p.char {
		if secondIdx < len(p.password) {
			return p.password[secondIdx] != p.char
		}
		return true
	}

	if secondIdx < len(p.password) {
		return p.password[secondIdx] == p.char
	}

	return false
}

func main() {
	var (
		f                        *os.File
		err                      error
		scanner                  *bufio.Scanner
		validCount, validV2Count int
		password                 *passwordPolicy
	)

	f, err = os.Open("02-input.txt")
	if err != nil {
		panic(err)
	}

	scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		if password, err = parsePasswordPolicy(scanner.Text()); err != nil {
			log.Fatalf("Failed to parse %s as password policy: %v", scanner.Text(), err)
		}

		if password.isValid() {
			validCount++
		}
		if password.isValidV2() {
			validV2Count++
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	fmt.Printf("V1 valid passwords: %d V2 valid passwords: %d\n", validCount, validV2Count)

}
