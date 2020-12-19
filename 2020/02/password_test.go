package main

import (
	"log"
	"testing"
)

func TestValidPolicyParsing(t *testing.T) {
	p, err := parsePasswordPolicy("1-3 a: abcde")

	if err != nil {
		log.Printf("Error expected to be nil, got %v", err)
		t.FailNow()
	}
	if p == nil {
		log.Printf("Expecting not-nill policy")
		t.FailNow()
	}

	if p.min != 1 {
		log.Printf("Repition min: expecting 1, got %d", p.min)
		t.FailNow()
	}
	if p.max != 3 {
		log.Printf("Repition max: expecting 3, got: %d ", p.max)
		t.FailNow()
	}

	if p.char != 'a' {
		log.Printf("Target char: expecting 'a', got %c", p.char)
		t.FailNow()
	}

	if p.password != "abcde" {
		log.Printf("Password: expecting 'abcde', got %s", p.password)
		t.FailNow()
	}

}

func TestValidPassword(t *testing.T) {
	var p = passwordPolicy{
		min:      1,
		max:      3,
		char:     'a',
		password: "abcde",
	}

	if p.isValid() == false {
		log.Printf("Expecting %v to be valid policy & password", p)
		t.FailNow()
	}
}

func TestInvalidPassword(t *testing.T) {
	var p = passwordPolicy{
		min:      1,
		max:      3,
		char:     'b',
		password: "cdefg",
	}

	if p.isValid() == true {
		log.Printf("Expecting %v to be **invalid** policy & password", p)
		t.FailNow()
	}
}

func TestSecondSchemeValidPassword(t *testing.T) {
	var p = passwordPolicy{
		min:      1,
		max:      3,
		char:     'a',
		password: "abcde",
	}

	if p.isValidV2() == false {
		log.Printf("Expecting %v to be valid v2 policy & password", p)
		t.FailNow()
	}
}

func TestSecondSchemeInvalidPassword(t *testing.T) {
	var p = passwordPolicy{
		min:      2,
		max:      9,
		char:     'c',
		password: "ccccccccc",
	}

	if p.isValidV2() == true {
		log.Printf("Expecting %v to be invalid v2 policy & password", p)
		t.FailNow()
	}
}
