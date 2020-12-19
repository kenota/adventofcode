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

func TestMapSubstractEmptySet(t *testing.T) {
	var (
		from = map[int]string{
			5: "",
			6: "",
		}
		what = map[int]string{}
		res  = map[int]string{}
		ok   bool
	)

	res = setSubstract(from, what)
	assert(len(res) == len(from), "No elements should be removed", t)

	_, ok = res[5]
	assert(ok, "5 should be present", t)
	_, ok = res[6]
	assert(ok, "6 Should be present", t)

}

func TestMapSubstractBiggerMap(t *testing.T) {
	var (
		from = map[int]string{
			5: "",
		}
		what = map[int]string{
			6: "",
			7: "",
			5: "",
		}
		res = map[int]string{}
	)

	res = setSubstract(from, what)
	assert(len(res) == 0, "Result should be empty", t)

}

func TestMapSubstractElements(t *testing.T) {
	var (
		from = map[int]string{
			5: "",
			6: "",
			7: "",
		}
		what = map[int]string{
			5: "",
			7: "",
		}
		res = map[int]string{}
		ok  bool
	)

	res = setSubstract(from, what)
	assert(len(res) == 1, "Only one element should be left", t)

	_, ok = res[6]
	assert(ok, "6 should be present", t)

}

func TestParseCorrectPassport(t *testing.T) {
	const passport = `hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm
`
	var (
		fields map[int]string
		ok     bool
		value  string
	)

	fields = parsePassportFields(passport)

	assert(len(fields) == 7, "passport should have 7 fields", t)

	value, ok = fields[HEIGHT]
	assert(ok, "height must be present", t)
	assert(value == "179cm", fmt.Sprintf("height should be 179cm, got: '%s'", value), t)

	value, ok = fields[ISSUE_YEAR]
	assert(ok, "issue year must be present", t)
	assert(value == "2013", fmt.Sprintf("issue year should be 2013, got: '%s'", value), t)

	value, ok = fields[EXPIRATION_YEAR]
	assert(ok, "expiration year must be present", t)
	assert(value == "2024", fmt.Sprintf("issue year should be 2024, got: '%s'", value), t)

	value, ok = fields[EYE_COLOR]
	assert(ok, "eye color year must be present", t)
	assert(value == "brn", fmt.Sprintf("issue year should be brn, got: '%s'", value), t)

	value, ok = fields[PASSPORT_ID]
	assert(ok, "passport id should be present", t)
	assert(value == "760753108", fmt.Sprintf("passport id should be brn, got: '%s'", value), t)

	value, ok = fields[BIRTHDAY_YEAR]
	assert(ok, "birthday year should be present", t)
	assert(value == "1931", fmt.Sprintf("birthday year should be 1931, got: '%s'", value), t)

	value, ok = fields[HAIR_COLOR]
	assert(ok, "hair color should be present", t)
	assert(value == "#ae17e1", fmt.Sprintf("hair colour should be 1931, got: '%s'", value), t)

}

func TestParseMultiplePassports(t *testing.T) {
	const data = `cid:123
eyr:2024

hgt:12
`
	var (
		parsed        []map[int]string
		first, second map[int]string
		ok            bool
	)

	parsed = parsePasswords(data)

	assert(len(parsed) == 2, "two passports should be parsed", t)

	first, second = parsed[0], parsed[1]

	assert(len(first) == 2, "first passport should have two fields", t)
	_, ok = first[COUNTRY_ID]
	assert(ok, "First passport should have country id field", t)
	_, ok = first[EXPIRATION_YEAR]
	assert(ok, "First passport should have an expiration year field", t)

	assert(len(second) == 1, "second passport should have one field", t)
	_, ok = second[HEIGHT]
	assert(ok, "second passport should have height field", t)
}

func TestIsValidPwOnValidPassword(t *testing.T) {
	const passport = `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
	byr:1937 iyr:2017 cid:147 hgt:183cm
`
	var p map[int]string

	p = parsePassportFields(passport)
	assert(isValidPassport(p), fmt.Sprintf("%v should be valid password", p), t)
}

func TestInvalidPassport(t *testing.T) {
	const passport = `hcl:#cfa07d eyr:2025 pid:166559648
	iyr:2011 ecl:brn hgt:59in
`
	var p map[int]string

	p = parsePassportFields(passport)
	assert(!isValidPassport(p), fmt.Sprintf("%v should be invalid password", p), t)
}

func TestPassportMissingCid(t *testing.T) {
	const passport = `hcl:#ae17e1 iyr:2013
	eyr:2024
	ecl:brn pid:760753108 byr:1931
	hgt:179cm`
	var p map[int]string

	p = parsePassportFields(passport)
	assert(isValidPassport(p), "Password with missing cid should be valid", t)
}

func TestValidBirthdayYears(t *testing.T) {
	var (
		years = []string{
			"1920", "1981", "2002",
		}
		year string
	)
	for _, year = range years {
		assert(isValidBirtdayYear(year), fmt.Sprintf("%s should be valid birthday year", year), t)
	}
}

func TestInvalidBirthdayYears(t *testing.T) {
	var (
		years = []string{
			"1919", "01", "20010", "2003",
		}
		year string
	)
	for _, year = range years {
		assert(!isValidBirtdayYear(year), fmt.Sprintf("%s should be invalid birthday year", year), t)
	}
}

func TestValidIssueYears(t *testing.T) {
	var (
		years = []string{
			"2010", "2011", "2012", "2015", "2020",
		}
		year string
	)
	for _, year = range years {
		assert(isValidIssueYear(year), fmt.Sprintf("%s should be valid issue year", year), t)
	}
}

func TestInvalidIssueYears(t *testing.T) {
	var (
		years = []string{
			"2009", "123", "0", "2021",
		}
		year string
	)
	for _, year = range years {
		assert(!isValidIssueYear(year), fmt.Sprintf("%s should be invalid issue year", year), t)
	}
}

/*
eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
*/

func TestValidExpirationYears(t *testing.T) {
	var (
		years = []string{
			"2020", "2025", "2030",
		}
		year string
	)
	for _, year = range years {
		assert(isValidExpirationYear(year), fmt.Sprintf("%s should be valid issue year", year), t)
	}
}

func TestInvalidExpirationYears(t *testing.T) {
	var (
		years = []string{
			"2019", "123", "0", "2031",
		}
		year string
	)
	for _, year = range years {
		assert(!isValidExpirationYear(year), fmt.Sprintf("%s should be invalid issue year", year), t)
	}
}

/*
hgt (Height) - a number followed by either cm or in:
If cm, the number must be at least 150 and at most 193.
If in, the number must be at least 59 and at most 76.
*/
func TestValidHeights(t *testing.T) {
	var (
		heights = []string{
			"150cm", "165cm", "193cm",
			"59in", "65in", "76in",
		}
		height string
	)
	for _, height = range heights {
		assert(isValidHeight(height), fmt.Sprintf("%s should be valid height", height), t)
	}
}

func TestInvalidHeights(t *testing.T) {
	var (
		heights = []string{
			"149cm", "155", "194cm",
			"58in", "999in", "77in",
		}
		height string
	)
	for _, height = range heights {
		assert(!isValidHeight(height), fmt.Sprintf("%s should be invalid height", height), t)
	}
}

func TestValidHairColors(t *testing.T) {
	var (
		tests = []string{
			"#123456", "#abcdef", "#abc123",
		}
		s string
	)
	for _, s = range tests {
		assert(isValidHairColor(s), fmt.Sprintf("%s should be valid hair color", s), t)
	}
}

func TestInvalidHairColors(t *testing.T) {
	var (
		tests = []string{
			"#12345g", "abcdef0", "#123456a",
		}
		s string
	)
	for _, s = range tests {
		assert(!isValidHairColor(s), fmt.Sprintf("%s should be invalid hair color", s), t)
	}
}

func TestValidEyeColors(t *testing.T) {
	var (
		tests = []string{
			"amb", "blu", "brn",
			"gry", "grn", "hzl",
			"oth",
		}
		s string
	)
	for _, s = range tests {
		assert(isValidEyeColor(s), fmt.Sprintf("%s should be valid eye color", s), t)
	}
}

func TestInvalidEyeColors(t *testing.T) {
	var (
		tests = []string{
			"amd", "ble", "#000fff",
		}
		s string
	)
	for _, s = range tests {
		assert(!isValidEyeColor(s), fmt.Sprintf("%s should be invalid eye color", s), t)
	}
}

func TestValidPassportIds(t *testing.T) {
	var (
		tests = []string{
			"123456789", "000000000", "000111111",
		}
		s string
	)
	for _, s = range tests {
		assert(isValidPassportID(s), fmt.Sprintf("%s should be valid passport id", s), t)
	}
}

func TestInvalidPassportIds(t *testing.T) {
	var (
		tests = []string{
			"12345678", "12345678A", "1234567890",
		}
		s string
	)
	for _, s = range tests {
		assert(!isValidPassportID(s), fmt.Sprintf("%s should be invalid eye color", s), t)
	}
}

/*
ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
pid (Passport ID) - a nine-digit number, including leading zeroes.
*/
