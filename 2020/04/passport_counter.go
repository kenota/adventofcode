package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	BIRTHDAY_YEAR = iota
	ISSUE_YEAR
	EXPIRATION_YEAR
	HEIGHT
	HAIR_COLOR
	EYE_COLOR
	PASSPORT_ID
	COUNTRY_ID
)

var REQUIRED_FIELDS = map[int]string{
	BIRTHDAY_YEAR:   "",
	ISSUE_YEAR:      "",
	EXPIRATION_YEAR: "",
	HEIGHT:          "",
	HAIR_COLOR:      "",
	EYE_COLOR:       "",
	PASSPORT_ID:     "",
}

func parsePassportFields(passport string) map[int]string {
	var (
		tokens         []string
		s, code, value string
		res            = map[int]string{}
	)

	tokens = strings.Fields(passport) // Tokens will be strings of format "cod:value"
	for _, s = range tokens[:] {
		code, value = s[0:3], s[4:]
		switch code {
		case "byr":
			res[BIRTHDAY_YEAR] = value
		case "iyr":
			res[ISSUE_YEAR] = value
		case "eyr":
			res[EXPIRATION_YEAR] = value
		case "hgt":
			res[HEIGHT] = value
		case "hcl":
			res[HAIR_COLOR] = value
		case "ecl":
			res[EYE_COLOR] = value
		case "pid":
			res[PASSPORT_ID] = value
		case "cid":
			res[COUNTRY_ID] = value
		default:
			panic(fmt.Sprintf("Uknown passport field: %s unparsed prefix: %s", code, s))
		}

	}

	return res
}

func parsePasswords(s string) []map[int]string {
	var (
		res     []map[int]string
		tokens  []string
		rawData string
	)

	tokens = strings.Split(s, "\n\n")
	for _, rawData = range tokens {
		res = append(res, parsePassportFields(rawData))
	}

	return res
}

func isValidPassport(p map[int]string) bool {
	return len(setSubstract(REQUIRED_FIELDS, p)) == 0
}

func setSubstract(from, what map[int]string) map[int]string {
	var (
		key   int
		value string
		res   = map[int]string{}
	)
	for key, value = range from {
		res[key] = value
	}
	for key = range what {
		delete(res, key)
	}

	return res
}

func parseFourDigitNumber(s string) (int, bool) {
	var (
		n   int64
		err error
	)

	if len(s) != 4 {
		return 0, false
	}

	if n, err = strconv.ParseInt(s, 10, 32); err != nil {
		return 0, false
	}

	return int(n), true
}

// byr (Birth Year) - four digits; at least 1920 and at most 2002.
func isValidBirtdayYear(value string) bool {
	var (
		year int
		ok   bool
	)

	if year, ok = parseFourDigitNumber(value); !ok {
		return false
	}

	return year >= 1920 && year <= 2002
}

// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
func isValidIssueYear(value string) bool {
	var (
		year int
		ok   bool
	)

	if year, ok = parseFourDigitNumber(value); !ok {
		return false
	}

	return year >= 2010 && year <= 2020
}

// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
func isValidExpirationYear(value string) bool {
	var (
		year int
		ok   bool
	)

	if year, ok = parseFourDigitNumber(value); !ok {
		return false
	}

	return year >= 2020 && year <= 2030
}

// hgt (Height) - a number followed by either cm or in:
// If cm, the number must be at least 150 and at most 193.
// If in, the number must be at least 59 and at most 76.
func isValidHeight(value string) bool {
	var (
		n   int64
		err error
	)

	if len(value) < 4 {
		return false
	}

	if n, err = strconv.ParseInt(value[:len(value)-2], 10, 32); err != nil {
		return false
	}
	switch value[len(value)-2:] {
	case "in":
		return n >= 59 && n <= 76
	case "cm":
		return n >= 150 && n <= 193
	default:
		return false
	}
}

func isValidHairColor(value string) bool {
	var c rune

	if len(value) != 7 {
		return false
	}
	if value[0] != '#' {
		return false
	}

	for _, c = range value[1:] {
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f')) {

			return false
		}
	}
	return true
}

// ecl (Eye Color) - exactly one of:  amb blu brn gry grn hzl oth.
func isValidEyeColor(value string) bool {
	return value == "amb" || value == "blu" || value == "brn" ||
		value == "gry" || value == "grn" || value == "hzl" || value == "oth"
}

// pid (Passport ID) - a nine-digit number, including leading zeroes.
func isValidPassportID(value string) bool {
	var (
		c rune
	)
	if len(value) != 9 {
		return false
	}

	for _, c = range value {
		if !(c >= '0' && c <= '9') {
			return false
		}
	}

	return true
}

func isValidFieldValue(field int, value string) bool {
	return false
}

func arePasswordFieldsValid(p map[int]string) bool {
	var (
		k int
		v string
	)

	for k, v = range p {
		var validator func(string) bool

		switch k {
		case BIRTHDAY_YEAR:
			validator = isValidBirtdayYear
		case ISSUE_YEAR:
			validator = isValidIssueYear
		case EXPIRATION_YEAR:
			validator = isValidExpirationYear
		case HEIGHT:
			validator = isValidHeight
		case HAIR_COLOR:
			validator = isValidHairColor
		case EYE_COLOR:
			validator = isValidEyeColor
		case PASSPORT_ID:
			validator = isValidPassportID
		case COUNTRY_ID:
			// no validation is needed for country id
			validator = func(string) bool {
				return true
			}
		}

		if !validator(v) {
			return false
		}
	}

	return true
}

func main() {
	var (
		err                     error
		data                    []byte
		passport                map[int]string
		f                       *os.File
		validStructurePassports int
		validPassports          int
	)

	if f, err = os.Open("04-input.txt"); err != nil {
		panic(err)
	}
	defer f.Close()

	if data, err = ioutil.ReadAll(f); err != nil {
		panic(err)
	}

	for _, passport = range parsePasswords(string(data)) {
		if isValidPassport(passport) {
			validStructurePassports++
			if arePasswordFieldsValid(passport) {
				validPassports++
			}
		}

	}

	fmt.Printf("Number of passports with valid structure: %d\n", validStructurePassports)
	fmt.Printf("Number of passports with valid structure: %d\n", validPassports)

}
